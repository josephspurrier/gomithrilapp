package component

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"app/api/model"
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

// DB .
type DB struct {
	Users []User `json:"users"`
}

// User .
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

	for _, v := range db.Users {
		if strings.ToLower(v.Username) == strings.ToLower(data.Username) {
			m.Body.Status = "user already exists"
			m.Body.Success = false
			// Send the response.
			b, _ := json.Marshal(m.Body)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, string(b))
			return
		}
	}

	db.Users = append(db.Users, User{
		Username: data.Username,
		Password: data.Password,
	})

	bb, err := json.Marshal(db)
	if err != nil {
		m.Body.Status = err.Error()
		m.Body.Success = false
		// Send the response.
		b, _ := json.Marshal(m.Body)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(b))
		return
	}

	err = ioutil.WriteFile("db.json", bb, 0644)
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
