package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/yuin/goldmark"
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

func gitAddAndCommit(notePath string, message string) error {
	// Get the absolute file path
	absPath, err := filepath.Abs(notePath)
	if err != nil {
		return fmt.Errorf("could not resolve absolute path: %w", err)
	}

	// Get the directory containing the file
	dir := filepath.Dir(absPath)

	// Find the git root (we assume the repo is at or above the notes folder)
	gitRoot := findGitRoot(dir)
	if gitRoot == "" {
		return fmt.Errorf("not inside a git repository")
	}

	// Run git commands from the root
	relPath, _ := filepath.Rel(gitRoot, absPath)

	addCmd := exec.Command("git", "add", relPath)
	addCmd.Dir = gitRoot
	if output, err := addCmd.CombinedOutput(); err != nil {
		fmt.Println("git add failed:", string(output))
		return err
	}

	commitCmd := exec.Command("git", "commit", "-m", message)
	commitCmd.Dir = gitRoot
	if output, err := commitCmd.CombinedOutput(); err != nil {
		// It's okay if there's nothing to commit
		if string(output) == "" || !isNoChangesError(string(output)) {
			fmt.Println("git commit failed:", string(output))
			return err
		}
	}

	return nil
}

func isNoChangesError(output string) bool {
	return output == "nothing to commit, working tree clean\n" ||
		output == "On branch main\nnothing to commit, working tree clean\n"
}

func findGitRoot(startDir string) string {
	dir := startDir
	for {
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}
		dir = parent
	}
}

func renderMarkdown(input []byte) []byte {
	var buf bytes.Buffer
	if err := goldmark.Convert(input, &buf); err != nil {
		return []byte("Markdown render error")
	}
	return buf.Bytes()
}
