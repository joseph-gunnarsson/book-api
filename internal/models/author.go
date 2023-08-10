package models

import (
	"github.com/joseph-gunnarsson/book-api/internal/database"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	FirstName   string `json:"firstName" gorm:"size:50;not null"`
	LastName    string `json:"lastName" gorm:"size:50;not null"`
	Nationality string `json:"nationality" gorm:"size:50;"`
	Website     string `json:"website" gorm:"size:50;"`
}

func CreateAuthor(author *Author) error {
	db := database.DB
	result := db.Create(author)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteAuthor(author *Author) error {
	db := database.DB
	result := db.Delete(author)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateAuthor(author *Author) error {
	db := database.DB
	result := db.Model(&author).Updates(author)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllAuthors() ([]Author, error) {
	db := database.DB
	var authors []Author
	result := db.Find(&authors)

	if result.Error != nil {
		return []Author{}, result.Error
	}

	return authors, nil
}

func GetAuthor(id uint) (Author, error) {
	db := database.DB
	var author Author
	result := db.First(&author, id)

	if result.Error != nil {
		return Author{}, result.Error
	}

	return author, nil
}

func GetAuthorByCondition(condition map[string]interface{}) ([]Author, error) {
	db := database.DB
	var authors []Author
	result := db.Where(condition).Find(&authors)

	if result.Error != nil {
		return []Author{}, result.Error
	}

	return authors, nil
}
