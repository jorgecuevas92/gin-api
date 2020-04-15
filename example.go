package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set the port
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	r := gin.Default()

	// Define Handlers
	rootHandler := func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	}

	pingHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}

	// Configure Routes
	r.GET("/", rootHandler)
	r.GET("/ping", pingHandler)

	// Run
	log.Printf("Listening on port %s", port)
	r.Run(":" + port)
}
