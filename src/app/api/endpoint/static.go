package endpoint

import (
	"errors"
	"fmt"
	"io/ioutil"
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

	p.Router.Get("/api/v1", p.Index)
	p.Router.Get("/static...", p.Static)
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
	if r.URL.Path == "/static/" {
		return http.StatusNotFound, nil
	}

	// Get the environment variable in production.
	basepath := os.Getenv("APP_ROOT")
	if len(basepath) == 0 {
		gopath := os.Getenv("GOPATH")
		if len(gopath) == 0 {
			return http.StatusInternalServerError, errors.New("could not find $APP_ROOT or $GOPATH environment variables")
		}

		basepath = filepath.Join(gopath, "src/app/ui/dist")
	}

	// If the file doesn't exist, serve the UI error message.
	fullPath := basepath + r.URL.Path
	if _, err := os.Stat(fullPath); err != nil {
		b, err := ioutil.ReadFile(basepath + "/index.html")
		if err != nil {
			return http.StatusInternalServerError, errors.New("could not find index.html")
		}

		// Return a 404 and serve the index.html file.
		w.WriteHeader(http.StatusNotFound)
		_, err = fmt.Fprint(w, string(b))
		if err != nil {
			return http.StatusInternalServerError, err
		}

		return 0, nil
	}

	// Serve the file to the user. Don't use filepath.join to protect against
	// "../" in the URL path.
	http.ServeFile(w, r, fullPath)

	return 0, nil
}
