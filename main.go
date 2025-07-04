package main

import (
	"html/template"

	"os"

	"github.com/gin-gonic/gin"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	if len(os.Args) > 1 {
		runCLI(os.Args[1:])
		return
	}

	runWebServer()
}

func runWebServer() {
	r := gin.Default()

	// Routes
	r.GET("/", HomeHandler)
	r.GET("/note/:name", NoteViewHandler)
	r.GET("/new", NewNoteFormHandler)
	r.POST("/new", NewNoteSubmitHandler)
	r.GET("/edit/:name", EditNoteFormHandler)
	r.POST("/edit/:name", EditNoteSubmitHandler)
	r.POST("/delete/:name", DeleteNoteHandler)

	r.Run(":8080")
}
