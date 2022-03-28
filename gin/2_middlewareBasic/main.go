package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func middleware(c *gin.Context) {
	fmt.Println("Hello Before")
	c.Next()
	fmt.Println("Hello After")
}

func main() {
	r := gin.Default()
	r.Use(middleware)
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("GET Ping")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
