package component

import (
	"encoding/json"
	"errors"
	"net/http"

	"app/api/store"
)

// RegisterEndpoint .
type RegisterEndpoint struct {
	Core
}

// SetupRegister .
func SetupRegister(core Core) {
	p := new(RegisterEndpoint)
	p.Core = core

	p.Router.Post("/register", p.Register)
}

// RegisterRequest is the request object.
type RegisterRequest struct {
	// Username for login.
	Username string `json:"username"`
	// Password for login.
	Password string `json:"password"`
}

// Register .
func (p *RegisterEndpoint) Register(w http.ResponseWriter, r *http.Request) (int, error) {
	// Convert the request to a struct.
	var data RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&data)
	defer r.Body.Close()
	if err != nil {
		return http.StatusBadRequest, err
	}

	user := store.NewUser(p.DB, p.Q)
	found, _, err := user.ExistsByField(user, "email", data.Username)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if found {
		return http.StatusBadRequest, errors.New("user already exists")
	}

	// Encrypt the password.
	password, err := p.Password.HashString(data.Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Create the user.
	ID, err := user.Create("first", "last", data.Username, password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return p.Response.Created(w, ID)
}
