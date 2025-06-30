# 📝 mdnote

**mdnote** is a simple, lightweight, Markdown-based note-taking web app built with Go, Gin, and Postgres (optionally). It uses plaintext `.md` files stored in a local `notes/` directory and supports Git-backed versioning to track edits over time.

## 🚀 Features

- Create, view, edit, and delete Markdown notes
- Notes are stored as `.md` files in a `notes/` directory
- Git integration to track changes and show last-edited timestamps
- HTML templating with `html/template`
- Minimal dependencies and easy to deploy

## 🧱 Tech Stack

- [Go](https://golang.org/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [blackfriday](https://github.com/russross/blackfriday) (Markdown parser)
- `html/template` for server-side rendering
- Optional: Git for version history

## 📁 Project Structure

```
mdnote/
├── main.go                # Application entry point
├── handlers.go            # Route handlers
├── utils.go               # Utility functions (e.g. git time, sanitization)
├── templates/             # HTML templates
│   ├── home.html
│   ├── note.html
│   ├── new.html
│   └── edit.html
├── notes/                 # Markdown notes directory
├── go.mod
├── go.sum
└── README.md
```

## ⚙️ Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/mdnote.git
cd mdnote
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Create the notes directory

```bash
mkdir notes
```

### 4. (Optional) Initialize Git for versioning

```bash
git init
git add notes
git commit -m "Initial commit of notes folder"
```

> **Important**: The app expects to find a `notes/` folder at the root level for all Markdown files.

### 5. Run the app

```bash
go run .
```

The app will be served at [http://localhost:8080](http://localhost:8080)

## ✍️ Usage

- **Home Page (`/`)**  
  Lists all notes with links to view each one.

- **View a Note (`/note/:name`)**  
  Renders Markdown as HTML. Shows last-edited time via Git if enabled.

- **Create a Note (`/new`)**  
  Fill out a form to add a new note.

- **Edit a Note (`/edit/:name`)**  
  Modify the Markdown content of an existing note.

- **Delete a Note (`/delete/:name`)**  
  Remove a note permanently. Requires confirmation.

## 🕒 Git-based Versioning (Optional)

To track the last edit time using Git:

1. Ensure your app is inside a Git repo (with `notes/` tracked).
2. The app uses `git log` to fetch the last commit timestamp for each `.md` file.
3. When editing/saving notes, the app can automatically:
    - `git add notes/note.md`
    - `git commit -m "Update note <title>"`

This allows you to see **“Last edited”** info on the note page.

## 🧪 Development Notes

- Templating is handled via `html/template` and parsed at startup.
- Markdown is rendered with [blackfriday](https://github.com/russross/blackfriday).
- Notes are stored in the file system — no database is required for the core app.

## 📦 Future Enhancements

- CLI for creating/editing notes from the terminal
- Git-backed diff viewer (view previous versions)
- Search or tag-based organization
- Database-backed storage (e.g., Postgres) as an alternative

## 📄 License

MIT © 2025 Lucas White
