package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

var (
	serverRoot       string // Absolute path of server root.
	staticFolderPath string // Absolute path of static file folder.
	faviconPath      string // Absolute path of "favicon.ico".
	indexTmplPath    string // Absolute path of index HTML template file.
)

// serveSingleFile serves Single Static File.
func serveSingleFile(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

// hello is home handler.
func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	t, _ := template.ParseFiles(indexTmplPath)
	d := struct {
		Title   string
		Content string
	}{
		"Hello, World!",
		"Static File Server Example:-)",
	}

	if err := t.Execute(w, &d); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetCurrentExecDir gets the current executable path.
// You may find more path helper functions in:
// https://github.com/northbright/pathhelper
func GetCurrentExecDir() (dir string, err error) {
	p, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}

	absPath, err := filepath.Abs(p)
	if err != nil {
		return "", err
	}

	dir = filepath.Dir(absPath)
	return dir, nil
}

func init() {
	// Get absolute path of server root(current executable).
	serverRoot, _ = GetCurrentExecDir()
	// Get static folder path.
	staticFolderPath = path.Join(serverRoot, "./static")
	// Get favicon.ico path.
	faviconPath = path.Join(serverRoot, "favicon.ico")
	// Get index template file path.
	indexTmplPath = path.Join(serverRoot, "./templates/index.tmpl")
}

func main() {
	// Index Handler
	http.HandleFunc("/", hello)

	// Serve Single File(Ex: favicon.ico)
	serveSingleFile("/favicon.ico", faviconPath)

	// Serve Static Files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticFolderPath))))

	fmt.Printf("Starting server...\nstatic folder abs path: %v\n", staticFolderPath)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
