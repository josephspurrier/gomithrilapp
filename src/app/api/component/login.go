package component

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"app/api/model"
	"app/api/pkg/webtoken"
	"app/api/store"
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
func (p *LoginEndpoint) Login(w http.ResponseWriter, r *http.Request) (int, error) {
	// Convert the request to a struct.
	var data LoginRequest

	err := json.NewDecoder(r.Body).Decode(&data)
	defer r.Body.Close()
	if err != nil {
		return http.StatusBadRequest, err
	}

	// Create the response.
	m := new(model.LoginResponse)
	m.Body.Success = false
	m.Body.Token = ""

	// Determine if the user exists.
	user := store.NewUser(p.DB, p.Q)
	found, ID, err := user.ExistsByField(user, "email", data.Username)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if !found {
		return http.StatusBadRequest, errors.New("user not found (1)")
	}

	// Populate the user.
	user.FindOneByID(user, ID)

	// Ensure the user's password matches.
	if !p.Password.MatchString(user.Password, data.Password) {
		return http.StatusBadRequest, errors.New("user not found (2)")
	}

	m.Body.Status = http.StatusText(http.StatusOK)
	m.Body.Success = true

	// Generate the access tokens.
	privateKey := []byte("asdfasdfasdf")
	t := new(webtoken.JWTAuth)
	t.Clock = webtoken.Clock{}
	t.PrivateKey = &privateKey
	u := new(webtoken.User)
	u.ID = ID
	at, _, err := t.GenerateTokens(u)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	m.Body.Token = at.Token
	log.Println(at.Token)

	return p.Response.JSON(w, m.Body)
}
