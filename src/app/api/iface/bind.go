package iface

import "net/http"

// IBind provides bind and validation for requests.
type IBind interface {
	Unmarshal(i interface{}, r *http.Request) (err error)
	Validate(s interface{}) error
}
