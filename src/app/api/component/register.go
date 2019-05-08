package component

import (
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

// Register .
func (p *RegisterEndpoint) Register(w http.ResponseWriter, r *http.Request) (int, error) {
	// swagger:parameters AuthLogin
	type request struct {
		// Required: true
		FirstName string `json:"first_name" validate:"required"`
		// in: formData
		// Required: true
		LastName string `json:"last_name" validate:"required"`
		// in: formData
		// Required: true
		Email string `json:"email" validate:"required,email"`
		// in: formData
		// Required: true
		Password string `json:"password" validate:"required"`
	}

	// Request validation.
	req := new(request)
	if err := p.Bind.JSONUnmarshal(req, r); err != nil {
		return http.StatusBadRequest, err
	} else if err = p.Bind.Validate(req); err != nil {
		return http.StatusBadRequest, err
	}

	user := store.NewUser(p.DB, p.Q)
	found, _, err := user.ExistsByField(user, "email", req.Email)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if found {
		return http.StatusBadRequest, errors.New("user already exists")
	}

	// Encrypt the password.
	password, err := p.Password.HashString(req.Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Create the user.
	ID, err := user.Create(req.FirstName, req.LastName, req.Email, password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return p.Response.Created(w, ID)
}
