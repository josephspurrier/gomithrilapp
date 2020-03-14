package endpoint

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// StaticEndpoint .
type StaticEndpoint struct {
	Core
}

// SetupStatic .
func SetupStatic(core Core) {
	p := new(StaticEndpoint)
	p.Core = core

	p.Router.Get("/api/v1", p.Index)
	p.Router.Get("/api/static...", p.Static)
}

// Index .
// swagger:route GET /api/v1 healthcheck Ready
//
// API is ready.
//
// Responses:
//   200: OKResponse
func (p StaticEndpoint) Index(w http.ResponseWriter, r *http.Request) (int, error) {
	return p.Response.OK(w, "ready")
}

// Static .
func (p StaticEndpoint) Static(w http.ResponseWriter, r *http.Request) (int, error) {
	if r.URL.Path == "/api/static/" {
		return http.StatusNotFound, nil
	}

	// Get the location of the executable.
	basepath, err := os.Executable()
	if err != nil {
		return http.StatusInternalServerError, nil
	}

	// If static folder is found to the executable, serve the file.
	staticPath := filepath.Join(basepath, "static")
	if stat, err := os.Stat(staticPath); err == nil && stat.IsDir() {
		// The static directory is found.
	} else if len(os.Getenv("GOPATH")) > 0 {
		// Else get the GOPATH.
		basepath = filepath.Join(os.Getenv("GOPATH"), "src/app/api")
	}

	// Serve the file to the user.
	http.ServeFile(w, r, filepath.Join(basepath, strings.TrimPrefix(r.URL.Path, "/api/")))

	return http.StatusOK, nil
}
