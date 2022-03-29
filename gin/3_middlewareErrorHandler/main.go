package main

import (
	"fmt"
	. "github.com/fullaheadcoco/go-examples/gin/3_middlewareErrorHandler/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	NotFoundError = fmt.Errorf("resource could not be found")
)

// https://golangexample.com/gin-wsderror-handling-middleware-is-a-middleware-for-the-popular-gin-framework/amp/
func main() {
	r := gin.Default()
	r.Use(
		ErrorHandler(
			Map(NotFoundError).ToResponse(func(c *gin.Context, err error) {
				c.Status(http.StatusNotFound)
				c.Writer.Write([]byte(err.Error()))
			}),
		),
	)

	r.GET("/ping", func(c *gin.Context) {
		c.Error(NotFoundError)
	})

	r.Run()
}
