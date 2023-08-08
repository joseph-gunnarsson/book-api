package routers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/joseph-gunnarsson/book-api/internal/models"
)

func AuthorRoutes(r *chi.Mux) {
	r.Get("/authors", GetAllAuthors)
	r.Post("/authors", CreateAuthor)
	r.Put("/authors/{id}", UpdateAuthor)
	r.Get("/authors/{id}", GetAuthorByID)
	r.Delete("/authors/{id}", DeleteAuthor)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	authorID, err := strconv.Atoi(id)
	if err != nil {
		handleErrorResponse(w, "Invalid author ID parameter", err, http.StatusBadRequest)
		return
	}

	author, err := models.GetAuthor(uint(authorID))
	if err != nil {
		handleErrorResponse(w, "Failed to get author", err, http.StatusInternalServerError)
		return
	}

	err = models.DeleteAuthor(&author)
	if err != nil {
		handleErrorResponse(w, "Failed to delete author", err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	responseJSON := map[string]interface{}{
		"message": "Author deleted successfully",
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

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	authorID, err := strconv.Atoi(id)
	if err != nil {
		handleErrorResponse(w, "Invalid author ID parameter", err, http.StatusBadRequest)
		return
	}

	var author models.Author
	err = json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		handleErrorResponse(w, "Failed to decode JSON", err, http.StatusBadRequest)
		return
	}
	author.ID = uint(authorID)

	err = models.UpdateAuthor(&author)
	if err != nil {
		handleErrorResponse(w, "Failed to update author", err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	responseJSON := map[string]interface{}{
		"message": "Author updated successfully",
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

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := models.GetAllAuthors()
	if err != nil {
		handleErrorResponse(w, "Failed to get authors", err, http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(authors)
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

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	err := json.NewDecoder(r.Body).Decode(&author)

	if err != nil {
		handleErrorResponse(w, "Failed to decode JSON", err, http.StatusBadRequest)
		return
	}
	err = models.CreateAuthor(&author)

	if err != nil {
		handleErrorResponse(w, "Failed to create author", err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	responseJSON := map[string]interface{}{
		"message": "Author created successfully",
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

func GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	authorID, err := strconv.Atoi(id)
	if err != nil {
		handleErrorResponse(w, "Invalid author ID parameter", err, http.StatusBadRequest)
		return
	}

	author, err := models.GetAuthor(uint(authorID))
	if err != nil {
		handleErrorResponse(w, "Failed to get author", err, http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(author)
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
