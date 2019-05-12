package testrequest

import (
	"io"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"app/api/component"
)

// SendForm is a helper to quickly make a form request.
func SendForm(t *testing.T, core component.Core, method string, target string,
	v url.Values) *httptest.ResponseRecorder {
	component.LoadRoutes(core)

	var body io.Reader
	if v != nil {
		body = strings.NewReader(v.Encode())
	}

	r := httptest.NewRequest(method, target, body)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	core.Router.ServeHTTP(w, r)

	return w
}
