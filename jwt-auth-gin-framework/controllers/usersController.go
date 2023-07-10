
package controllers

import (
	"fmt"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"jwt-auth/initializers"
	"jwt-auth/models"
	"github.com/golang-jwt/jwt/v5"
)


func Signup(c *gin.Context){
	var body struct  {
		Fullname string
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body!",
		})
		return 
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10, )


	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password!",
		})
		return 
	}

	user := models.User{
		Email: body.Email,
		Password: string(hash),
		Fullname: body.Fullname,
	}

	fmt.Println(body)

	result := initializers.DB.Create(&user)

	if result.Error != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create User!",
		})
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registration success!",
	})

}


func Login(c *gin.Context){
	var body struct  {
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body!",
		})
		return 
	}
	var user models.User

	initializers.DB.First(&user, "email = ?" , body.Email);



	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password!",
		})
		return 
	}


	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password!",
		})
		return 
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,	
		"exp": time.Now().Add(time.Hour * 24 * 30  ).Unix(),
	})

	tokenString, err := token.SignedString([]byte("3424234234"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create JWT!",
		})
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Success!",
		"token": tokenString,
	})

}


func UserDetails(c *gin.Context){

	user, _  := c.Get(
		"user",
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "User Details!",
		"user": user,
	})
}

func AllUsers(c *gin.Context){
	var users []models.User
	result :=initializers.DB.Find(&users)


	if result.Error != nil {
         c.JSON( http.StatusInternalServerError, gin.H{ "Message": "Query failed",
		   "error": result.Error,
		})
		return
    }
	c.JSON(http.StatusOK, gin.H{
		"message": "User Details!",
		"users": users,
	})

}