package component

import (
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

// Login .
func (p *LoginEndpoint) Login(w http.ResponseWriter, r *http.Request) (int, error) {
	// swagger:parameters AuthLogin
	type request struct {
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

	// Determine if the user exists.
	user := store.NewUser(p.DB, p.Q)
	found, ID, err := user.ExistsByField(user, "email", req.Email)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if !found {
		return http.StatusBadRequest, errors.New("user not found (1)")
	}

	// Populate the user.
	user.FindOneByID(user, ID)

	// Ensure the user's password matches.
	if !p.Password.MatchString(user.Password, req.Password) {
		return http.StatusBadRequest, errors.New("user not found (2)")
	}

	// Create the response.
	m := new(model.LoginResponse)
	m.Body.Token = ""
	m.Body.Status = http.StatusText(http.StatusOK)

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
