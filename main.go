package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// Routes
	r.GET("/", HomeHandler)
	r.GET("/note/:name", NoteViewHandler)
	r.GET("/new", NewNoteFormHandler)
	r.POST("/new", NewNoteSubmitHandler)
	r.GET("/edit/:name", EditNoteFormHandler)
	r.POST("/edit/:name", EditNoteSubmitHandler)

	r.Run(":8080")
}
