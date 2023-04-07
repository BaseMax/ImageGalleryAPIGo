package handlers

import (
	"fmt"
	"net/http"

	"github.com/BaseMax/ImageGalleryAPIGo/utils"
)

func UploadImgWithMetadata(w http.ResponseWriter, r *http.Request) {
	utils.Logger(r)
	fmt.Fprintf(w, "this is UploadImgWithMetadata")
}

func GetAllImages(w http.ResponseWriter, r *http.Request) {
	utils.Logger(r)
	fmt.Fprintf(w, "this is GetAllImages")
}

func GetOneImgWithMetadata(w http.ResponseWriter, r *http.Request) {
	utils.Logger(r)
	fmt.Fprintf(w, "this is GetOneImgWithMetadata")
}

func UpdateImgMetadata(w http.ResponseWriter, r *http.Request) {
	utils.Logger(r)
	fmt.Fprintf(w, "this is UpdateImgMetadata")
}

func DeleteImg(w http.ResponseWriter, r *http.Request) {
	utils.Logger(r)
	fmt.Fprintf(w, "this is DeleteImg")
}
