package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/joseph-gunnarsson/book-api/internal/database"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string    `json:"title" gorm:"size:255;not null;unique;"`
	ReleaseDate time.Time `json:"releaseDate" gorm:"not null"`
	Genre       []Genre   `gorm:"many2many:book_genre;"`
	Description string    `json:"description" gorm:"size:1000"`
	ISBN        string    `json:"isbn" gorm:"size:12;not null"`
	AuthorID    int       `gorm:"index;not null" json:"authorID"`
	Author      Author    `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE;" json:"author"`
}

func CreateBook(book *Book) error {
	db := database.DB

	if err := ValidateGenreIDs(book.Genre); err != nil {
		return err
	}

	result := db.Omit("Author").Create(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteBookByID(bookID uint) error {
	db := database.DB

	var existingBook Book
	if err := db.First(&existingBook, bookID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("book with ID %d does not exist", bookID)
		}
		return err
	}

	result := db.Delete(&existingBook)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func ValidateGenreIDs(genres []Genre) error {
	db := database.DB

	for _, genre := range genres {
		var existingGenre Genre
		if err := db.First(&existingGenre, genre.ID).Error; err != nil {
			return fmt.Errorf("genre with ID %d doesn't exist", genre.ID)
		}
	}
	return nil
}

func UpdateBook(book *Book) error {
	db := database.DB

	if err := ValidateGenreIDs(book.Genre); err != nil {
		return err
	}

	err := db.Model(&book).Association("Genre").Replace(book.Genre)
	if err != nil {
		return err
	}

	result := db.Model(&book).Updates(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllBooks() ([]Book, error) {
	db := database.DB
	var books []Book
	result := db.Preload("Author").Preload("Genre").Find(&books)

	if result.Error != nil {
		return []Book{}, result.Error
	}
	return books, nil
}

func GetBookById(id int) (Book, error) {
	db := database.DB
	var book Book
	result := db.Preload("Author").Preload("Genre").First(&book, id)

	if result.Error != nil {
		return Book{}, result.Error
	}

	return book, nil
}

func GetBookByCondition(condition map[string]interface{}) ([]Book, error) {
	db := database.DB
	var books []Book
	result := db.Preload("Author").Preload("Genre").Where(condition).Find(&books)

	if result.Error != nil {
		return []Book{}, result.Error
	}

	return books, nil
}
