package main

import (
	"html/template"
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

	data := struct {
		Notes []string
	}{
		Notes: notes,
	}

	c.Header("Content-Type", "text/html")
	if err := templates.ExecuteTemplate(c.Writer, "home.html", data); err != nil {
		c.String(http.StatusInternalServerError, "Template rendering error: %v", err)
	}
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
	lastEdited, err := getGitLastEditedTime(path)
	if err != nil {
		lastEdited = "unknown"
	}

	data := struct {
		Title      string
		Content    template.HTML
		LastEdited string
	}{
		Title:      name,
		Content:    template.HTML(htmlContent),
		LastEdited: lastEdited,
	}

	c.Header("Content-Type", "text/html")
	if err := templates.ExecuteTemplate(c.Writer, "note.html", data); err != nil {
		c.String(http.StatusInternalServerError, "Template rendering error: %v", err)
	}
}

func NewNoteFormHandler(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	if err := templates.ExecuteTemplate(c.Writer, "new.html", nil); err != nil {
		c.String(http.StatusInternalServerError, "Template rendering error: %v", err)
	}
}

func NewNoteSubmitHandler(c *gin.Context) {
	title := sanitizeFileName(c.PostForm("title"))
	content := c.PostForm("content")

	if title == "" {
		c.String(http.StatusBadRequest, "Invalid title")
		return
	}

	path := filepath.Join("notes", title+".md")
	if _, err := os.Stat(path); err == nil {
		c.String(http.StatusBadRequest, "Note already exists")
		return
	}

	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error saving note")
		return
	}

	c.Redirect(http.StatusFound, "/note/"+title)
}

func EditNoteFormHandler(c *gin.Context) {
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

	data := struct {
		Title   string
		Content string
	}{
		Title:   name,
		Content: string(content),
	}

	c.Header("Content-Type", "text/html")
	if err := templates.ExecuteTemplate(c.Writer, "edit.html", data); err != nil {
		c.String(http.StatusInternalServerError, "Template rendering error: %v", err)
	}
}

func EditNoteSubmitHandler(c *gin.Context) {
	name := sanitizeFileName(c.Param("name"))
	if name == "" {
		c.String(http.StatusBadRequest, "Invalid note name")
		return
	}

	content := c.PostForm("content")
	path := filepath.Join("notes", name+".md")

	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error saving note")
		return
	}

	c.Redirect(http.StatusFound, "/note/"+name)
}

func DeleteNoteHandler(c *gin.Context) {
	name := sanitizeFileName(c.Param("name"))
	if name == "" {
		c.String(http.StatusBadRequest, "Invalid note name")
		return
	}

	path := filepath.Join("notes", name+".md")
	err := os.Remove(path)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error deleting note")
		return
	}

	c.Redirect(http.StatusFound, "/")
}
