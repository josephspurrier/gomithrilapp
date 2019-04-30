package component

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"app/api/model"
	"app/api/pkg/webtoken"
)

// LoginEndpoint .
type LoginEndpoint struct {
	Core
}

// SetupLogin .
func SetupLogin(core Core) {
	p := new(LoginEndpoint)
	p.Core = core

	p.Router.Post("/login", p.Login)
}

// LoginRequest is the request object.
type LoginRequest struct {
	// Username for login.
	Username string `json:"username"`
	// Password for login.
	Password string `json:"password"`
}

// Login .
func (p *LoginEndpoint) Login(w http.ResponseWriter, r *http.Request) {
	// Convert the request to a struct.
	var data LoginRequest

	err := json.NewDecoder(r.Body).Decode(&data)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Error in request: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Create the response.
	m := new(model.LoginResponse)
	m.Body.Success = false
	m.Body.Token = ""

	dbFile, err := ioutil.ReadFile("db.json")
	if err != nil {
		m.Body.Status = err.Error()
		m.Body.Success = false
		// Send the response.
		b, _ := json.Marshal(m.Body)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(b))
		return
	}

	db := new(DB)

	err = json.Unmarshal(dbFile, db)
	if err != nil {
		m.Body.Status = err.Error()
		m.Body.Success = false
		// Send the response.
		b, _ := json.Marshal(m.Body)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(b))
		return
	}

	userFound := false

	for _, v := range db.Users {
		if strings.ToLower(v.Username) == strings.ToLower(data.Username) &&
			v.Password == data.Password {
			userFound = true
			break
		}
	}

	if !userFound {
		m.Body.Status = "access denied"
		m.Body.Success = false
		// Send the response.
		b, _ := json.Marshal(m.Body)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(b))
		return
	}

	m.Body.Status = http.StatusText(http.StatusOK)
	m.Body.Success = true

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

	m.Body.Token = at.Token
	log.Println(at.Token)

	// Send the response.
	b, _ := json.Marshal(m.Body)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(b))
}
