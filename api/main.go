package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/josephspurrier/govueapp/api/webtoken"
)

func init() {
	// Verbose logging with file name and line number.
	log.SetFlags(log.Lshortfile)
}

func main() {
	http.Handle("/", Handler(http.HandlerFunc(Index)))

	log.Println("Server started.")
	http.ListenAndServe(":8081", nil)
}

// Login is the request object.
type Login struct {
	// Username for login.
	Username string `json:"username"`
	// Password for login.
	Password string `json:"password"`
}

// LoginResponse is the response object.
type LoginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

// Index handles the initial page logic.
func Index(w http.ResponseWriter, r *http.Request) {
	switch {
	// Serve the index file.
	case r.URL.Path == "/":
		http.ServeFile(w, r, "public/index.html")
		return
	// Serve files out of the static file.
	case strings.HasPrefix(r.URL.Path, "/static/"):
		path := "public/" + r.URL.Path[1:]

		// Only serve files.
		if fi, err := os.Stat(path); err == nil && !fi.IsDir() {
			http.ServeFile(w, r, path)
			return
		}
	case r.URL.Path == "/data":
		// Allow browser to access API endpoint.
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// If a preflight request, then respond with the allowed headers.
		if r.Method == "OPTIONS" {
			// Allow browser to send headers for preflight request.
			w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
			return
		}

		// Convert the request to a struct.
		var data Login

		if r.Method != http.MethodPost {
			http.Error(w, "Request must be a: "+http.MethodPost, http.StatusBadRequest)
			return
		}

		err := json.NewDecoder(r.Body).Decode(&data)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, "Error in request: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Create the response.
		m := LoginResponse{
			Success: false,
			Token:   "",
		}

		// Send an invalid message back.
		if data.Username != "a" || data.Password != "a" {
			// Send the response.
			b, _ := json.Marshal(m)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, string(b))
			return
		}

		m.Success = true

		// Generate the access tokens.
		privateKey := []byte("asdfasdfasdf")
		t := new(webtoken.JWTAuth)
		t.Clock = webtoken.Clock{}
		t.PrivateKey = &privateKey
		u := new(webtoken.User)
		u.ID = 12
		at, _, err := t.GenerateTokens(u)
		if err != nil {
			log.Println(err)
			return
		}

		m.Token = at.Token
		log.Println(at.Token)

		// Send the response.
		b, _ := json.Marshal(m)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(b))
		return
	}

	http.NotFound(w, r)
}

// Handler will log the HTTP requests.
func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
