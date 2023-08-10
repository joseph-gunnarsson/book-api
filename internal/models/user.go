package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"size:30;not null;unique"`
	Password string `json:"-" gorm:"size:72;not null;"`
}

func CreateUser(user *User) {

}
