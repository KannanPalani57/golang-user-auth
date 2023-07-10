package models

import (
	"github.com/KannanPalani57/go-mux-user-auth/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Fullname string `gorm:"size:191" json:"fullname,omitempty"`
	Email    string `gorm:"unique" json:"email"`
	Password string ` json:"password"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	db.Create(&u)
	return u
}

// func GetAllBooks() []User {
// 	var Users []User
// 	db.Find(&Users)
// 	return Users

// }
