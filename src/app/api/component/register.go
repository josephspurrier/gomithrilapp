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

	p.Router.Post("/v1/register", p.Register)
}

// Register .
// swagger:route POST /v1/register user UserRegister
//
// Register a new user.
//
// Security:
//   token:
//
// Responses:
//   201: CreatedResponse
//   400: BadRequestResponse
//   401: UnauthorizedResponse
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
	fullRequest := new(request)
	req := fullRequest.Body
	if err := p.Bind.JSONUnmarshal(fullRequest, r); err != nil {
		return http.StatusBadRequest, err
	} else if err = p.Bind.Validate(fullRequest); err != nil {
		return http.StatusBadRequest, err
	}

	// Determine if the user already exists.
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
