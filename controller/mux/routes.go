package mux

import "net/http"

func (r *rest) routing() {
	r.router = r.router.StrictSlash(true)
	api := r.router.PathPrefix("/api").Subrouter()
	{
		// Get all images
		api.HandleFunc("/images/", r.handler.getAll).Methods(http.MethodGet)
		// Upload a new image
		api.HandleFunc("/images/", r.handler.upload).Methods(http.MethodPost)
		// Retrieve a single image with its metadata
		api.HandleFunc("/images/{_id}", r.handler.getOne).Methods(http.MethodGet)
		// Update the metadata for a single image.
		api.HandleFunc("/images/{_id}", r.handler.update).Methods(http.MethodPut)
		// Delete a single image.
		api.HandleFunc("/images/{_id}", r.handler.delete).Methods(http.MethodDelete)
	}
}
