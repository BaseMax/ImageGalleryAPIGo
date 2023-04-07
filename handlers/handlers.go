package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
)

func UploadImgWithMetadata(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "this is UploadImgWithMetadata")
	}
}

func GetAllImages(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "this is GetAllImages")
	}
}

func GetOneImgWithMetadata(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "this is GetOneImgWithMetadata")
	}
}

func UpdateImgMetadata(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "this is UpdateImgMetadata")
	}
}

func DeleteImg(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "this is DeleteImg")
	}
}
