package ui

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Handler will handle any UI requests.
func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") || strings.HasPrefix(r.URL.Path, "/static/") {
			next.ServeHTTP(w, r)
		} else {
			UI(w, r)
		}
	})
}

// UI will serve the UI files.
func UI(w http.ResponseWriter, r *http.Request) {
	// Get the location of the executable.
	basepath, err := os.Executable()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// If static folder is found to the executable, serve the file.
	staticPath := filepath.Join(basepath, "static")
	if stat, err := os.Stat(staticPath); err == nil && stat.IsDir() {
		// The static directory is found.
	} else if len(os.Getenv("GOPATH")) > 0 {
		// Else get the GOPATH.
		basepath = filepath.Join(os.Getenv("GOPATH"), "src/app/ui")
	}

	// Serve the index file.
	http.ServeFile(w, r, filepath.Join(basepath, "index.html"))
}
