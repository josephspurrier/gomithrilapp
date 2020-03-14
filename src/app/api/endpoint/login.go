package endpoint

import (
	"errors"
	"net/http"

	"app/api/model"
)

// LoginEndpoint .
type LoginEndpoint struct {
	Core
}

// SetupLogin .
func SetupLogin(c Core) {
	p := new(LoginEndpoint)
	p.Core = c

	p.Router.Post("/api/v1/login", p.Login)
}

// Login .
// swagger:route POST /api/v1/login authentication UserLogin
//
// Authenticate a user.
//
// Responses:
//   200: LoginResponse
//   400: BadRequestResponse
//   500: InternalServerErrorResponse
func (p *LoginEndpoint) Login(w http.ResponseWriter, r *http.Request) (int, error) {
	// swagger:parameters UserLogin
	type request struct {
		// in: body
		Body struct {
			// Required: true
			Email string `json:"email" validate:"required,email"`
			// Required: true
			Password string `json:"password" validate:"required"`
		}
	}

	// Request validation.
	req := new(request)
	if err := p.Bind.UnmarshalAndValidate(req, r); err != nil {
		return http.StatusBadRequest, err
	}

	// Determine if the user exists.
	user := p.Store.User.New()
	found, err := p.Store.User.FindOneByField(&user, "email", req.Body.Email)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if !found {
		return http.StatusBadRequest, errors.New("login information does not match")
	}

	// Ensure the user's password matches. Use the same error message to prevent
	// brute-force from finding usernames.
	if !p.Password.Match(user.Password, req.Body.Password) {
		return http.StatusBadRequest, errors.New("login information does not match")
	}

	// Create the response.
	m := new(model.LoginResponse).Body
	m.Status = http.StatusText(http.StatusOK)

	// Generate the access token.
	m.Token, err = p.Token.Generate(user.ID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return p.Response.JSON(w, m)
}
