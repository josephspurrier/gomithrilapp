package endpoint

import (
	"errors"
	"net/http"
)

// RegisterEndpoint .
type RegisterEndpoint struct {
	Core
}

// SetupRegister .
func SetupRegister(core Core) {
	p := new(RegisterEndpoint)
	p.Core = core

	p.Router.Post("/v1/register", p.Register)
}

// Register .
// swagger:route POST /v1/register authentication UserRegister
//
// Register a user.
//
// Responses:
//   201: CreatedResponse
//   400: BadRequestResponse
//   500: InternalServerErrorResponse
func (p *RegisterEndpoint) Register(w http.ResponseWriter, r *http.Request) (int, error) {
	// swagger:parameters UserRegister
	type request struct {
		// in: body
		Body struct {
			// Required: true
			FirstName string `json:"first_name" validate:"required"`
			// Required: true
			LastName string `json:"last_name" validate:"required"`
			// Required: true
			Email string `json:"email" validate:"required,email"`
			// Required: true
			Password string `json:"password" validate:"required"`
		}
	}

	// Request validation.
	req := new(request)
	if err := p.Bind.Unmarshal(req, r); err != nil {
		return http.StatusBadRequest, err
	} else if err = p.Bind.Validate(req); err != nil {
		return http.StatusBadRequest, err
	}

	// Determine if the user already exists.
	user := p.Store.User.New()
	found, _, err := p.Store.User.ExistsByField(&user, "email", req.Body.Email)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if found {
		return http.StatusBadRequest, errors.New("user already exists")
	}

	// Encrypt the password.
	password, err := p.Password.Hash(req.Body.Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Create the user.
	ID, err := p.Store.User.Create(req.Body.FirstName,
		req.Body.LastName, req.Body.Email, password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return p.Response.Created(w, ID)
}
