package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/BaseMax/ImageGalleryAPIGo/utils"
)

func UploadImgWithMetadata(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.Logger(r)
		fmt.Fprintf(w, "this is UploadImgWithMetadata")
	}
}

func GetAllImages(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.Logger(r)
		fmt.Fprintf(w, "this is GetAllImages")
	}
}

func GetOneImgWithMetadata(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.Logger(r)
		fmt.Fprintf(w, "this is GetOneImgWithMetadata")
	}
}

func UpdateImgMetadata(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.Logger(r)
		fmt.Fprintf(w, "this is UpdateImgMetadata")
	}
}

func DeleteImg(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.Logger(r)
		fmt.Fprintf(w, "this is DeleteImg")
	}
}
