package main

import (
	"io/ioutil"
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

func NewNoteFormHandler(c *gin.Context) {
	form := `
    <html><head><title>New Note</title></head><body>
    <a href="/">Home</a>
    <h1>Create New Note</h1>
    <form action="/new" method="POST">
      Title:<br><input name="title" /><br><br>
      Content:<br><textarea name="content" rows="20" cols="80"></textarea><br><br>
      <input type="submit" value="Save" />
    </form>
    </body></html>
    `
	c.Header("Content-Type", "text/html")
	c.String(200, form)
}

func NewNoteSubmitHandler(c *gin.Context) {
	title := sanitizeFileName(c.PostForm("title"))
	content := c.PostForm("content")

	if title == "" {
		c.String(400, "Invalid title")
		return
	}

	path := filepath.Join("notes", title+".md")
	if _, err := os.Stat(path); err == nil {
		c.String(400, "Note already exists")
		return
	}

	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		c.String(500, "Error saving note")
		return
	}

	c.Redirect(302, "/note/"+title)
}

func EditNoteFormHandler(c *gin.Context) {
	name := sanitizeFileName(c.Param("name"))
	if name == "" {
		c.String(400, "Invalid note name")
		return
	}

	path := filepath.Join("notes", name+".md")
	content, err := ioutil.ReadFile(path)
	if err != nil {
		c.String(404, "Note not found")
		return
	}

	form := `
    <html><head><title>Edit Note</title></head><body>
    <a href="/">Home</a> | <a href="/note/` + name + `">View Note</a>
    <h1>Edit Note: ` + name + `</h1>
    <form action="/edit/` + name + `" method="POST">
      <textarea name="content" rows="20" cols="80">` + string(content) + `</textarea><br><br>
      <input type="submit" value="Save" />
    </form>
    </body></html>
    `
	c.Header("Content-Type", "text/html")
	c.String(200, form)
}

func EditNoteSubmitHandler(c *gin.Context) {
	name := sanitizeFileName(c.Param("name"))
	if name == "" {
		c.String(400, "Invalid note name")
		return
	}

	content := c.PostForm("content")
	path := filepath.Join("notes", name+".md")

	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		c.String(500, "Error saving note")
		return
	}

	c.Redirect(302, "/note/"+name)
}
