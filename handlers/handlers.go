package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/BaseMax/ImageGalleryAPIGo/utils"
	"github.com/gorilla/mux"
)

type ImageMetadata struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
	FilePath    string `json:"file_path"`
	CreatedAt   string `json:"created_at"`
}

func UploadImgWithMetadata(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.Logger(r)
		// Parse multipart form data
		err := r.ParseMultipartForm(10 << 20) // 10 MB
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Get image file from form data
		file, handler, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Get other metadata from form data
		title := r.FormValue("title")
		description := r.FormValue("description")
		tags := r.FormValue("tags")

		// Generate file name
		fileExt := filepath.Ext(handler.Filename)
		fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), fileExt)

		// Create file on file system
		fileDir := "uploads/"
		if _, err := os.Stat(fileDir); os.IsNotExist(err) {
			os.Mkdir(fileDir, os.ModePerm)
		}
		filePath := fileDir + fileName
		f, err := os.Create(filePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// Copy image file to file system
		_, err = io.Copy(f, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Insert metadata into database
		createdAt := time.Now().Format("2006-01-02 15:04:05")
		stmt, err := db.Prepare("INSERT INTO images(title, description, tags, file_path, created_at) VALUES(?, ?, ?, ?, ?)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		res, err := stmt.Exec(title, description, tags, filePath, createdAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Get the ID of the inserted image
		imageID, err := res.LastInsertId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Return the inserted image metadata as JSON
		image := ImageMetadata{
			ID:          int(imageID),
			Title:       title,
			Description: description,
			Tags:        tags,
			FilePath:    filePath,
			CreatedAt:   createdAt,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(image)
	}
}

func GetAllImages(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.Logger(r)
		// Query the database for all images
		rows, err := db.Query("SELECT id, title, description, tags, file_path, created_at FROM images")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Loop through the rows of the result set and create an array of Image structs
		var images []ImageMetadata
		for rows.Next() {
			var image ImageMetadata
			err := rows.Scan(&image.ID, &image.Title, &image.Description, &image.Tags, &image.FilePath, &image.CreatedAt)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			images = append(images, image)
		}

		// Marshal the array of Image structs to JSON and return it in the response
		imageJSON, err := json.Marshal(images)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(imageJSON)
	}
}

func GetOneImgWithMetadata(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.Logger(r)
		// Parse the ID from the URL parameter
		vars := mux.Vars(r)
		id := vars["id"]

		// Query the database for the image with the specified ID
		row := db.QueryRow("SELECT id, title, description, tags, file_path, created_at FROM images WHERE id = ?", id)

		// Scan the result row into an Image struct
		var image ImageMetadata
		err := row.Scan(&image.ID, &image.Title, &image.Description, &image.Tags, &image.FilePath, &image.CreatedAt)
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Marshal the Image struct to JSON and return it in the response
		imageJSON, err := json.Marshal(image)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(imageJSON)
	}
}

func UpdateImgMetadata(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.Logger(r)
		// Parse the ID from the URL parameter
		vars := mux.Vars(r)
		id := vars["id"]

		// Decode the request body into an Image struct
		var image ImageMetadata
		err := json.NewDecoder(r.Body).Decode(&image)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Update the metadata of the image with the specified ID in the database
		_, err = db.Exec("UPDATE images SET title = ?, description = ?, tags = ? WHERE id = ?", image.Title, image.Description, image.Tags, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Return a success response
		imageJSON, err := json.Marshal(image)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(imageJSON)
	}
}

func DeleteImg(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.Logger(r)
		// Parse the image ID from the request URL
		vars := mux.Vars(r)
		imageID := vars["id"]

		// Get the image metadata from the database
		var image ImageMetadata
		err := db.QueryRow("SELECT id, title, description, tags, file_path, created_at FROM images WHERE id = ?", imageID).Scan(&image.ID, &image.Title, &image.Description, &image.Tags, &image.FilePath, &image.CreatedAt)
		if err == sql.ErrNoRows {
			http.Error(w, fmt.Sprintf("Image with ID %s not found", imageID), http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Delete the image file from the filesystem
		err = os.Remove(image.FilePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Delete the image metadata from the database
		_, err = db.Exec("DELETE FROM images WHERE id = ?", imageID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Return a success response
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Image with ID %s has been deleted", imageID)
	}
}
