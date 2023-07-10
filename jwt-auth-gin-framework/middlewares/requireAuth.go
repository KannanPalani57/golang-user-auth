package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"jwt-auth/initializers"
	"jwt-auth/models"
)

func RequireAuth(c *gin.Context){
	tokenString := c.GetHeader("x-access-token")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
	
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("3424234234"), nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)

	}
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {


	if	float64(time.Now().Unix()) > claims["exp"].(float64){
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	var user models.User

	initializers.DB.First(&user, claims["sub"])

	if user.ID == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

		c.Set("user", user)

	
		c.Next()
		} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}


}
