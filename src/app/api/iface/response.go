package iface

import "net/http"

// IResponse provides outputs for data.
type IResponse interface {
	JSON(w http.ResponseWriter, body interface{}) (int, error)
	Created(w http.ResponseWriter, recordID string) (int, error)
	OK(w http.ResponseWriter, message string) (int, error)
}
