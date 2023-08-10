package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joseph-gunnarsson/book-api/internal/database"
	"github.com/joseph-gunnarsson/book-api/internal/models"
	"github.com/joseph-gunnarsson/book-api/internal/routers"
)

func main() {
	err := database.InitDB()

	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	// Drop and recreate tables
	err = database.DB.Migrator().DropTable(&models.Author{}, &models.Book{}, &models.Genre{})
	if err != nil {
		log.Fatal("Failed to drop tables:", err)
	}
	database.DB.AutoMigrate(&models.Author{}, &models.Book{}, &models.Genre{})

	// Insert dummy data
	author := models.Author{
		FirstName:   "J.K.",
		LastName:    "Rowling",
		Nationality: "British",
		Website:     "https://www.jkrowling.com/",
	}
	err = models.CreateAuthor(&author)
	if err != nil {
		log.Println("Failed to create author:", err)
	}

	genre1 := models.Genre{
		Genre: "Fantasy",
	}
	err = models.CreateGenre(&genre1)
	if err != nil {
		log.Println("Failed to create genre:", err)
	}

	genre2 := models.Genre{
		Genre: "Adventure",
	}
	err = models.CreateGenre(&genre2)
	if err != nil {
		log.Println("Failed to create genre:", err)
	}

	book := models.Book{
		Title:       "Harry Potter and the Sorcerer's Stone",
		ReleaseDate: time.Date(1997, time.June, 26, 0, 0, 0, 0, time.UTC),
		Genre:       []models.Genre{genre1, genre2},
		Description: "The first book in the Harry Potter series.",
		ISBN:        "97805903427",
		AuthorID:    int(author.ID),
	}
	err = models.CreateBook(&book)
	if err != nil {
		log.Println("Failed to create book:", err)
	}
	r := chi.NewRouter()
	routers.BookRoutes(r)
	routers.GenreRoutes(r)
	routers.AuthorRoutes(r)

	port := 8080
	fmt.Printf("Server started on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
