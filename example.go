package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

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

	// Parameters in the path

	// This handler will match /users/john but will not match /users/ or /users

	r.GET("/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// This handler will match /user/john and /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john

	r.GET("/users/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		// For each matched request Context will hold the route definition
		hasAction := c.FullPath() == "/users/:name/*action"
		log.Printf("Route has action: %s", strconv.FormatBool(hasAction))
		c.String(http.StatusOK, message)
	})

	// Querystring parameters

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe

	r.GET("/welcome", func(c *gin.Context) {
		firstName := c.DefaultQuery("firstname", "Guest")
		lastName := c.Query("lastname") //shortcut for c.Request.URL.Query().Get("lastname")
		c.String(http.StatusOK, "Hello %s %s", firstName, lastName)
	})

	// Multipart/URLEncoded Form
	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	// Query + Post Form

	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		log.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	// Map as querystring or postform parameters

	r.POST("postmap", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		log.Printf("ids: %v; names: %v", ids, names)
	})

	// Run

	log.Printf("Listening on port %s", port)
	r.Run(":" + port)
}
