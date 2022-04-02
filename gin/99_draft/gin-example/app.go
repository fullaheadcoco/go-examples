package main

import (
	// Log items to the terminal
	"log"

	// Import gin for route definition
	"github.com/gin-gonic/gin"
	// Import godotenv for .env variables
	"github.com/joho/godotenv"
	// Import our app controllers
	"github.com/fullaheadcoco/gin-example/controllers"
)

// init gets called before the main function
func init() {
	// Log error if .env file does not exist
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

func main() {
	// Init gin router
	router := gin.Default()

	// It's great to version your API's
	v1 := router.Group("/api/v1")
	{
		// Define the hello controller
		hello := new(controllers.HelloWorldController)
		// Define a GET request to call the Default
		// method in controllers/hello.go
		v1.GET("/hello", hello.Default)

		// Define the user controller
		user := new(controllers.UserController)
		// Create the signup endpoint
		v1.POST("/signup", user.Signup)

	}

	// Handle error response when a route is not defined
	router.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(404, gin.H{"message": "Not found"})
	})

	// Init our server
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
