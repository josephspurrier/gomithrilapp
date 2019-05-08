package iface

import "net/http"

// IBind provides bind and validation for requests.
type IBind interface {
	FormUnmarshal(i interface{}, r *http.Request) (err error)
	JSONUnmarshal(i interface{}, r *http.Request) (err error)
	Validate(s interface{}) error
}
