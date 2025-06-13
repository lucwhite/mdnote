package main

import (
	"os"
	"os/exec"
	"strings"
	"time"
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

func getGitLastEditedTime(filePath string) (string, error) {
	cmd := exec.Command("git", "log", "-1", "--format=%ci", "--", filePath)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	dateStr := strings.TrimSpace(string(output))
	if dateStr == "" {
		return "untracked", nil
	}
	// Parse date string from git output (e.g. "2025-06-13 20:00:00 -0400")
	t, err := time.Parse("2006-01-02 15:04:05 -0700", dateStr)
	if err != nil {
		return dateStr, nil // fallback to raw string
	}
	return t.Format("Jan 2, 2006 15:04 MST"), nil
}
