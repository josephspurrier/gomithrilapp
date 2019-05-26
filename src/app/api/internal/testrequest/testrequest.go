package testrequest

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"app/api/boot"
	"app/api/endpoint"
)

// TR is a test request.
type TR struct {
	// Header clears after use.
	Header http.Header
	// SkipMiddleware resets to false after use.
	SkipMiddleware bool
	// SkipRoutes resets to false after use.
	SkipRoutes bool
}

// New returns a new TR.
func New() *TR {
	return &TR{
		Header:         make(http.Header),
		SkipMiddleware: false,
		SkipRoutes:     false,
	}
}

// SendForm is a helper to quickly make a form request.
func (tr *TR) SendForm(t *testing.T, core endpoint.Core, method string, target string,
	v url.Values) *httptest.ResponseRecorder {
	if !tr.SkipRoutes {
		// Load the routes.
		boot.LoadRoutes(core)
	}

	var body io.Reader
	if v != nil {
		body = strings.NewReader(v.Encode())
	}

	r := httptest.NewRequest(method, target, body)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range tr.Header {
		r.Header.Set(k, v[0])
	}
	tr.Header = make(http.Header)

	w := httptest.NewRecorder()

	if !tr.SkipMiddleware {
		boot.Middleware(core).ServeHTTP(w, r)
	} else {
		tr.SkipMiddleware = false
		core.Router.ServeHTTP(w, r)
	}

	return w
}

// SendJSON is a helper to quickly make a JSON request.
func (tr *TR) SendJSON(t *testing.T, core endpoint.Core, method string, target string,
	v url.Values) *httptest.ResponseRecorder {
	if !tr.SkipRoutes {
		// Load the routes.
		boot.LoadRoutes(core)
	}

	var body io.Reader
	if v != nil {
		body = strings.NewReader(ToJSON(v))
	}

	r := httptest.NewRequest(method, target, body)

	r.Header.Set("Content-Type", "application/json")
	for k, v := range tr.Header {
		r.Header.Set(k, v[0])
	}
	tr.Header = make(http.Header)

	w := httptest.NewRecorder()

	if !tr.SkipMiddleware {
		boot.Middleware(core).ServeHTTP(w, r)
	} else {
		tr.SkipMiddleware = false
		core.Router.ServeHTTP(w, r)
	}

	return w
}

// ToJSON converts a url.Values to a JSON string.
func ToJSON(values url.Values) string {
	m := make(map[string]string)

	for k, v := range values {
		if len(v) > 0 {
			m[k] = v[0]
		} else {
			m[k] = ""
		}
	}

	js, _ := json.Marshal(m)

	return string(js)
}
