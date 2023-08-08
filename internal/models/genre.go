package models

import (
	"github.com/joseph-gunnarsson/book-api/internal/database"
	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model
	Genre string `json:"genre" gorm:"size:255;not null;unique;"`
}

func CreateGenre(genre *Genre) error {
	db := database.DB
	result := db.Create(genre)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteGenre(genre *Genre) error {
	db := database.DB
	result := db.Delete(genre)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateGenre(genre *Genre) error {
	db := database.DB
	result := db.Save(genre)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetAllGenre() ([]Genre, error) {
	db := database.DB
	var genres []Genre
	result := db.Find(&genres)

	if result.Error != nil {
		return []Genre{}, result.Error
	}

	return genres, nil
}

func GetByName(name string) (Genre, error) {
	db := database.DB
	var genre Genre
	result := db.Where("genre = ?", name).First(&genre)

	if result.Error != nil {
		return Genre{}, result.Error
	}

	return genre, nil
}
