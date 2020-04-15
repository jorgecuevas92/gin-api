package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	r := gin.New()

	rootHandler := func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	}

	pingHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}

	r.GET("/", rootHandler)
	r.GET("/ping", pingHandler)
	log.Printf("Listening on port %s", port)
	r.Run(":" + port)
}
