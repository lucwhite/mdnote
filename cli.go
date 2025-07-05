package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func runCLI(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: mdnote [new|edit|view|delete|list]")
		return
	}

	command := args[0]
	var title string
	var editor string

	// Extract --editor flag
	for i, arg := range args {
		if arg == "--editor" && i+1 < len(args) {
			editor = args[i+1]
		}
	}

	// Grab the note title (first non-flag arg after command)
	for _, arg := range args[1:] {
		if !strings.HasPrefix(arg, "--") && title == "" {
			title = arg
		}
	}
	switch command {
	case "new":
		if title == "" {
			fmt.Println("Usage: mdnote new <title> [--editor editor]")
			return
		}
		createOrEditNote(title, false, editor)

	case "edit", "open": // alias
		if title == "" {
			fmt.Println("Usage: mdnote edit <title> [--editor editor]")
			return
		}
		createOrEditNote(title, true, editor)
	case "view":
		viewNote(args[1])
	case "delete":
		deleteNoteCLI(args[1])
	case "list":
		listNotesCLI()
	case "serve":
		runWebServer()
	case "update":
		updateGit(args[1])
	default:
		fmt.Println("Unknown command:", command)
	}
}

func createOrEditNote(name string, editing bool, editor string) {
	path := resolveNotePath(name)

	if _, err := os.Stat(path); os.IsNotExist(err) && editing {
		fmt.Println("Note not found.")
		return
	}

	if editor == "" {
		editor = os.Getenv("EDITOR")
		if editor == "" {
			editor = "subl"
		}
	}

	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to open editor (%s): %v\n", editor, err)
		return
	}
	// Changes are not yet made, so we cannot commit here.
	//_ = gitAddAndCommit(path, "Update note via CLI: "+name)
}

func viewNote(name string) {
	path := resolveNotePath(name)
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Note not found.")
		return
	}
	fmt.Println(string(content))
}

func deleteNoteCLI(name string) {
	path := resolveNotePath(name)
	err := os.Remove(path)
	if err != nil {
		fmt.Println("Error deleting note:", err)
		return
	}
	_ = gitAddAndCommit(path, "Delete note via CLI: "+name)
	fmt.Println("Deleted:", name)
}

func listNotesCLI() {
	notes, _ := ListNotes()
	for _, note := range notes {
		fmt.Println("-", note)
	}
}

func updateGit(noteName string) {
	// Ensure we’re pointing to the actual notes directory
	notePath := resolveNotePath(noteName)
	err := gitAddAndCommit(notePath, "Update note: "+noteName)
	if err != nil {
		fmt.Println("Git update failed:", err)
	}
}
