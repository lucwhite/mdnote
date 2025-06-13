package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func HomeHandler(c *gin.Context) {
	notes, err := ListNotes()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error reading notes: %v", err)
		return
	}

	html := "<h1>mdnote</h1><ul>"
	for _, note := range notes {
		html += `<li><a href="/note/` + note + `">` + note + `</a></li>`
	}
	html += "</ul><a href=\"/new\">Create New Note</a>"

	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, html)
}

func NoteViewHandler(c *gin.Context) {
	name := sanitizeFileName(c.Param("name"))
	if name == "" {
		c.String(http.StatusBadRequest, "Invalid note name")
		return
	}

	path := filepath.Join("notes", name+".md")
	content, err := os.ReadFile(path)
	if err != nil {
		c.String(http.StatusNotFound, "Note not found")
		return
	}

	htmlContent := blackfriday.Run(content)
	page := `
    <html><head><title>` + name + `</title></head><body>
    <a href="/">Home</a> | <a href="/edit/` + name + `">Edit</a><hr>`
	page += string(htmlContent)
	page += `</body></html>`

	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, page)
}
