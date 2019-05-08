package component

import (
	"encoding/json"
	"fmt"
	"net/http"

	"app/api/model"
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
func (p *RegisterEndpoint) Register(w http.ResponseWriter, r *http.Request) {
	// Convert the request to a struct.
	var data RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&data)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Error in request: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Create the response.
	m := new(model.RegisterResponse)
	m.Body.Success = false

	user := store.NewUser(p.DB, p.Q)
	found, _, err := user.ExistsByField(user, "email", data.Username)
	if err != nil {
		m.Body.Status = err.Error()
		m.Body.Success = false
		// Send the response.
		b, _ := json.Marshal(m.Body)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(b))
		return
	} else if found {
		m.Body.Status = "user already exists"
		m.Body.Success = false
		// Send the response.
		b, _ := json.Marshal(m.Body)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(b))
		return
	}

	// Encrypt the password.
	password, err := p.Password.HashString(data.Password)
	if err != nil {
		//return http.StatusInternalServerError, err
		m.Body.Status = err.Error()
		m.Body.Success = false
		// Send the response.
		b, _ := json.Marshal(m.Body)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(b))
		return
	}

	_, err = user.Create("first", "last", data.Username, password)
	if err != nil {
		m.Body.Status = err.Error()
		m.Body.Success = false
		// Send the response.
		b, _ := json.Marshal(m.Body)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(b))
		return
	}

	m.Body.Status = http.StatusText(http.StatusOK)
	m.Body.Success = true

	// Send the response.
	b, _ := json.Marshal(m.Body)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(b))
}
