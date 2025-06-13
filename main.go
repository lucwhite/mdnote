package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// Routes
	r.GET("/", HomeHandler)
	r.GET("/note/:name", NoteViewHandler)
	// Add more routes here later

	r.Run(":8080")
}
