package jwt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"app/api"
	"app/api/model"
)

// Config contains the dependencies for the handler.
type Config struct {
	whitelist []string
	webtoken  api.IToken
	ctx       api.IContext
}

// New returns a new loq request middleware.
func New(whitelist []string, webtoken api.IToken, ctx api.IContext) *Config {
	return &Config{
		whitelist: whitelist,
		ctx:       ctx,
		webtoken:  webtoken,
	}
}

// Handler will require a JWT.
func (c *Config) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Determine if the page is in the JWT whitelist.
		if !IsWhitelisted(r.Method, r.URL.Path, c.whitelist) {
			// Require JWT on all routes.
			bearer := r.Header.Get("Authorization")

			// If the token is missing, show an error.
			if len(bearer) < 8 || !strings.HasPrefix(bearer, "Bearer ") {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				r := new(model.UnauthorizedResponse)
				r.Body.Status = http.StatusText(http.StatusUnauthorized)
				r.Body.Message = "authorization token is missing"
				_ = json.NewEncoder(w).Encode(r.Body)
				return
			}

			userID, err := c.webtoken.Verify(bearer[7:])
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				r := new(model.UnauthorizedResponse)
				r.Body.Status = http.StatusText(http.StatusUnauthorized)
				r.Body.Message = "authorization token is invalid"
				_ = json.NewEncoder(w).Encode(r.Body)
				return
			}

			c.ctx.SetUserID(r, userID)
		}
		next.ServeHTTP(w, r)
	})
}

// IsWhitelisted returns true if the request is in the whitelist. If only an
// asterisk is found in the whitelist, allow all routes. If an asterisk is
// found in the page string, then whitelist only the matching paths.
func IsWhitelisted(method string, path string, arr []string) (found bool) {
	s := fmt.Sprintf("%v %v", method, path)
	for _, i := range arr {
		if i == "*" || s == i {
			return true
		} else if strings.Contains(i, "*") {
			if strings.HasPrefix(s, i[:strings.Index(i, "*")]) {
				return true
			}
		}
	}
	return
}
