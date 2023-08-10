package routers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joseph-gunnarsson/book-api/internal/models"
)

func GenreRoutes(r *chi.Mux) {
	r.Get("/genres", GetAllGenres)
	r.Post("/genres", CreateGenre)
	r.Get("/genres/{name}", getGenreByName)
	r.Put("/genres/{name}", UpdateGenre)
	r.Delete("/genres/{name}", DeleteGenre)
}

func getGenreByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	genre, err := models.GetGenreByName(name)
	if err != nil {
		handleErrorResponse(w, "Failed to get genre", err, http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(genre)

	if err != nil {
		handleErrorResponse(w, "Failed to marshal response", err, http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func DeleteGenre(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	genre, err := models.GetGenreByName(name)
	if err != nil {
		handleErrorResponse(w, "Failed to get genre", err, http.StatusBadRequest)
		return
	}

	err = models.DeleteGenre(&genre)
	if err != nil {
		handleErrorResponse(w, "Failed to delete genre", err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	responseJSON := map[string]interface{}{
		"message": "Genre deleted successfully",
	}

	data, err := json.Marshal(responseJSON)
	if err != nil {
		handleErrorResponse(w, "Failed to marshal response", err, http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func UpdateGenre(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	var genre models.Genre
	err := json.NewDecoder(r.Body).Decode(&genre)
	if err != nil {
		handleErrorResponse(w, "Failed to decode JSON", err, http.StatusBadRequest)
		return
	}
	genre.Genre = name

	err = models.UpdateGenre(&genre)
	if err != nil {
		handleErrorResponse(w, "Failed to update genre", err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	responseJSON := map[string]interface{}{
		"message": "Genre updated successfully",
	}

	data, err := json.Marshal(responseJSON)
	if err != nil {
		handleErrorResponse(w, "Failed to marshal response", err, http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func GetAllGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := models.GetAllGenre()
	if err != nil {
		handleErrorResponse(w, "Failed to get genres", err, http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(genres)
	if err != nil {
		handleErrorResponse(w, "Failed to marshal data", err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		handleErrorResponse(w, "Failed to write response", err, http.StatusInternalServerError)
		return
	}
}

func CreateGenre(w http.ResponseWriter, r *http.Request) {
	var genre models.Genre
	err := json.NewDecoder(r.Body).Decode(&genre)

	if err != nil {
		handleErrorResponse(w, "Failed to decode JSON", err, http.StatusBadRequest)
		return
	}
	err = models.CreateGenre(&genre)

	if err != nil {
		handleErrorResponse(w, "Failed to create genre", err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	responseJSON := map[string]interface{}{
		"message": "Genre created successfully",
	}

	data, err := json.Marshal(responseJSON)
	if err != nil {
		handleErrorResponse(w, "Failed to marshal response", err, http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
