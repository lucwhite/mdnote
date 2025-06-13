package main

import (
	"os"
	"strings"
)

func sanitizeFileName(name string) string {
	filtered := ""
	for _, ch := range name {
		if (ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') ||
			ch == '-' || ch == '_' {
			filtered += string(ch)
		}
	}
	return filtered
}

func ListNotes() ([]string, error) {
	files, err := os.ReadDir("./notes")
	if err != nil {
		return nil, err
	}

	var notes []string
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".md") {
			name := strings.TrimSuffix(f.Name(), ".md")
			notes = append(notes, name)
		}
	}
	return notes, nil
}
