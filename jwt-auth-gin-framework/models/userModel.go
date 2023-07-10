package models

import (
	"gorm.io/gorm"

)

type User struct {
	gorm.Model
	Fullname string `gorm:"unique"`
	Email string `gorm:"unique"`
	Password string
}