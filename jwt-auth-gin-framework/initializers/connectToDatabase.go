package initializers

import (
	// "os"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jwt-auth/models"

)

var DB 	*gorm.DB;

func ConnectToDB(){
	var err error
	dsn := "root:test@tcp(localhost:3306)/user_auth?parseTime=true"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect DB")
	}

	fmt.Println(DB);
}


func SyncDatabase(){
	DB.AutoMigrate(&models.User{})
}