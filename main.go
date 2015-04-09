package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Server is ready")
	})
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "waiting for implementation", "status": http.StatusOK})
	})
	r.GET("/users/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "waiting for implementation", "status": http.StatusOK})
	})

	r.POST("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "waiting for implementation", "status": http.StatusOK})
	})

	r.POST("/users/auth", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "waiting for implementation", "status": http.StatusOK})
	})

	log.Fatal(r.Run(":8080"))
}
