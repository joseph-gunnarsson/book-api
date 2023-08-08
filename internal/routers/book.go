package routers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/joseph-gunnarsson/book-api/internal/models"
)

func BookRoutes(r *chi.Mux) {

	r.Get("/books", GetAllBooks)
	r.Post("/books", CreateBook)
	r.Put("/books/{id}", UpdateBook)
	r.Get("/books/{id}", GetBookById)
	r.Delete("/books/{id}", DeleteBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		handleErrorResponse(w, "Invalid book ID parameter", err, http.StatusBadRequest)
		return
	}

	err = models.DeleteBookByID(uint(bookID))

	if err != nil {
		handleErrorResponse(w, "Failed to delete book", err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	responseJSON := map[string]interface{}{
		"message": "Book deleted successfully",
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

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		handleErrorResponse(w, "Invalid book ID parameter", err, http.StatusBadRequest)
		return
	}

	var book models.Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		handleErrorResponse(w, "Failed to decode JSON", err, http.StatusBadRequest)
		return
	}
	book.ID = uint(bookID)

	err = models.UpdateBook(&book)
	if err != nil {
		handleErrorResponse(w, "Failed to update book", err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	responseJSON := map[string]interface{}{
		"message": "Book updated successfully",
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

func handleErrorResponse(w http.ResponseWriter, errMsg string, err error, statusCode int) {
	log.Printf("%s: %v", errMsg, err)
	http.Error(w, errMsg, statusCode)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetAllBooks()
	if err != nil {
		handleErrorResponse(w, "Failed to get books", err, http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(books)
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

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		handleErrorResponse(w, "Failed to decode json", err, http.StatusBadRequest)
		return
	}
	err = models.CreateBook(&book)

	if err != nil {
		handleErrorResponse(w, "Failed to create book", err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	responseJSON := map[string]interface{}{
		"message": "Book created successfully",
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

func GetBookById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	book, err := models.GetBookById(id)
	if err != nil {
		handleErrorResponse(w, "Invalid book ID parameter", err, http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(book)
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
