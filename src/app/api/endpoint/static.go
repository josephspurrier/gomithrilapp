package endpoint

import (
	"net/http"
	"os"
	"path/filepath"
)

// StaticEndpoint .
type StaticEndpoint struct {
	Core
}

// SetupStatic .
func SetupStatic(core Core) {
	p := new(StaticEndpoint)
	p.Core = core

	p.Router.Get("/v1", p.Index)
	p.Router.Get("/static...", p.Static)
}

// Index .
// swagger:route GET /v1 healthcheck Ready
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
	if r.URL.Path == "/static/" {
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
	http.ServeFile(w, r, filepath.Join(basepath, r.URL.Path[1:]))

	return http.StatusOK, nil
}
