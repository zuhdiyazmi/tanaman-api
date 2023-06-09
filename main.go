package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zuhdiyazmi/go-tanaman/database"
	"github.com/zuhdiyazmi/go-tanaman/routes"
)

func main() {
	r := gin.Default()
	database.ConnectDatabase()

	routes.SetupTanamanRoutes(r)
	routes.SetupWishlistRoutes(r)

	r.Use(ErrorHandler)

	r.Run()
}

// ErrorHandler is a middleware to handle error response
func ErrorHandler(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"message": "internal server error",
			})
		}
	}()
	c.Next()
}
