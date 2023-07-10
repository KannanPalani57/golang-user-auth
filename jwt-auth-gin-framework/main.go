package main

import (
	// "os"///////////
	// "fmt"
	"jwt-auth/initializers"
	"jwt-auth/controllers"
	"github.com/gin-gonic/gin"
	"jwt-auth/middlewares"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func main(){
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/me", middlewares.RequireAuth,    controllers.UserDetails)   
	r.GET("/users",   controllers.AllUsers)   
	r.Run() 
	// fmt.Println(os.Getenv("PORT"))
}


