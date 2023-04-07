package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/BaseMax/ImageGalleryAPIGo/handlers"
	"github.com/BaseMax/ImageGalleryAPIGo/utils"
)

type Image struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
	CreatedAt   string `json:"created_at"`
}

func main() {
	// Initialize database
	if err := godotenv.Load("env.env"); err != nil {
		log.Fatal("Error while loading env.env file!")
	}

	envVars := []string{"DBUSER", "DBPASS", "NET", "ADDR", "DBNAME"}
	if _, err := utils.CheckEnvVars(envVars); err != nil {
		log.Fatal("We have some problem in env variables!")
	}

	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  os.Getenv("NET"),
		Addr:                 os.Getenv("ADDR"),
		DBName:               os.Getenv("DBNAME"),
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Println("Connected to MySQL!")

	defer db.Close()

	if _, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS images (
		id INT NOT NULL AUTO_INCREMENT,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		tags VARCHAR(255),
		created_at DATETIME NOT NULL,
		PRIMARY KEY (id)
	);`); err != nil {
		log.Fatal(err)
	}

	var wait time.Duration
	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/api/images").Subrouter()
	// Upload a new image and its metadata.
	api.HandleFunc("/", handlers.UploadImgWithMetadata(db)).Methods(http.MethodPost)

	// Retrieve a list of all images with their metadata.
	api.HandleFunc("/", handlers.GetAllImages(db)).Methods(http.MethodGet)

	// Retrieve a single image with its metadata.
	api.HandleFunc("/{id:[1-9]+}", handlers.GetOneImgWithMetadata(db)).Methods(http.MethodGet)

	// Update the metadata for a single image.
	api.HandleFunc("/{id:[1-9]+}", handlers.UpdateImgMetadata(db)).Methods(http.MethodPut)

	// Delete a single image.
	api.HandleFunc("/{id:[1-9]+}", handlers.DeleteImg(db)).Methods(http.MethodDelete)

	srv := &http.Server{
		Addr:         "127.0.0.1:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
