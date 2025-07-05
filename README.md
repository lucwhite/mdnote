# 📘 mdnote

`mdnote` is a simple Markdown-based note-taking app built with Go, Gin, and Markdown rendering. It supports a web UI and a CLI for creating, editing, and viewing notes stored as `.md` files.

---

## 🚀 Features

- Create, edit, delete, and view notes in your browser
- Notes are saved as `.md` files in a `notes/` directory
- Markdown is rendered using [Goldmark](https://github.com/yuin/goldmark)
- Git-backed history (optional)
- Templated HTML with minimal styling via [Water.css](https://watercss.kognise.dev/)
- CLI support for working with notes directly from the terminal

---

## 📦 Getting Started

### 1. Clone and build the app

```bash
git clone https://github.com/your-username/mdnote.git
cd mdnote
go build -o mdnote .
```

### 2. Run the web server

```bash
./mdnote serve
```
Visit http://localhost:8080 to view the app.

## 🧑‍💻 CLI Usage

You can also interact with mdnote via the command line.

### 🔧 Available Commands

```bash
mdnote new <title> [--editor editor]     # Create a new note (opens in editor)
mdnote edit <title> [--editor editor]    # Edit an existing note (alias: open)
mdnote open <title>                      # Alias for edit
mdnote view <title>                      # Print note content to stdout
mdnote delete <title>                    # Delete a note
mdnote list                              # List all note titles
mdnote update <title>                    # Commit changes to a note to git
mdnote serve                             # Start the web server
```

> Default editor is subl (Sublime Text). You can override with --editor code, --editor nano, etc., or set the EDITOR environment variable.

**Note:**  
You can pass note titles with or without the `.md` extension (e.g., `mdnote view mynote` or `mdnote view mynote.md` both work).

## Example Usage

```bash
mdnote new "dev-notes" --editor subl
mdnote edit "dev-notes"
mdnote view "dev-notes"
mdnote update "dev-notes"
mdnote list
mdnote serve
```

To build the CLI:

```bash
go build -o mdnote .
./mdnote list
```

To install it globally:
```bash
sudo mv mdnote /usr/local/bin/
```

Now you can run mdnote from anywhere in your terminal!

## 📄 License
MIT