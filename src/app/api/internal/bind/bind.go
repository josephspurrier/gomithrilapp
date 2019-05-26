package bind

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

	validator "gopkg.in/go-playground/validator.v9"
)

// IRouter extracts a URL parameter value.
type IRouter interface {
	Param(r *http.Request, param string) string
}

// Binder contains the request bind an validator objects.
type Binder struct {
	validator *validator.Validate
	router    IRouter
}

// New returns a new binder for request bind and validation.
func New(r IRouter) *Binder {
	return &Binder{
		validator: validator.New(),
		router:    r,
	}
}

// Validate will validate a struct using the validator.
func (b *Binder) Validate(s interface{}) error {
	return b.validator.Struct(s)
}

// Unmarshal will perform an unmarshal on an interface using: form or JSON.
func (b *Binder) Unmarshal(iface interface{}, r *http.Request) (err error) {
	// Check for errors.
	v := reflect.ValueOf(iface)
	if v.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer, not a value")
	}

	// Load the map.
	m := make(map[string]interface{})

	// Try to auto detect data type based on on the header.
	switch r.Header.Get("Content-Type") {
	case "", "application/x-www-form-urlencoded":
		// Parse the form.
		err = r.ParseForm()
		if err != nil {
			return err
		}

		for k, vv := range r.Form {
			m[k] = vv[0]
		}
	case "application/json":
		// Decode to the interface.
		err = json.NewDecoder(r.Body).Decode(&m)
		r.Body.Close()
		if err != nil {
			// No longer fail on an unmarshal error. This is so users can submit
			// empty data for GET requests, yet we can still map the URL
			// parameter by using the same logic.
		}

		// Copy the map items to a new map.
		mt := make(map[string]interface{})
		for key, value := range m {
			mt[key] = value
		}

		// Save the map to the body to handle cases where there is a body
		// defined.
		m["body"] = mt
	}

	// Loop through each field to extract the URL parameter.
	elem := reflect.Indirect(v.Elem())
	keys := elem.Type()
	for j := 0; j < elem.NumField(); j++ {
		tag := keys.Field(j).Tag
		tagvalue := tag.Get("json")
		pathParam := b.router.Param(r, tagvalue)
		if len(pathParam) > 0 {
			m[tagvalue] = pathParam
		}
	}

	// Convert to JSON.
	var data []byte
	data, err = json.Marshal(m)
	if err != nil {
		return
	}

	// Unmarshal to the interface from JSON.
	return json.Unmarshal(data, &iface)
}
