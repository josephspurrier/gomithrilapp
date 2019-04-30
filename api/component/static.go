package component

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/josephspurrier/govueapp/api/pkg/app"
)

// StaticEndpoint .
type StaticEndpoint struct {
	app.Core
}

// SetupStatic .
func SetupStatic(core app.Core) {
	p := new(StaticEndpoint)
	p.Core = core

	p.Router.Get("/", p.IndexGET)
	p.Router.Get("/static/*", p.StaticGET)
}

// IndexGET .
func (p StaticEndpoint) IndexGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"api": "ready"}`)
}

// StaticGET .
func (p StaticEndpoint) StaticGET(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, "/") {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, r.URL.Path[1:])
}
