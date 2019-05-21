package response

import (
	"encoding/json"
	"net/http"

	"app/api/model"
)

// Output is the response object.
type Output struct{}

// New returns a new response object.
func New() *Output {
	return &Output{}
}

// JSON will output JSON to the writer.
func (o *Output) JSON(w http.ResponseWriter, body interface{}) (int, error) {
	// Write the content.
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// OK will write an OK status to the writer.
func (o *Output) OK(w http.ResponseWriter, message string) (int, error) {
	r := new(model.OKResponse)
	r.Body.Status = http.StatusText(http.StatusOK)
	r.Body.Message = message

	// Write the content.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(r.Body)

	return http.StatusOK, nil
}

// Created will output a creation response to the writer.
func (o *Output) Created(w http.ResponseWriter, recordID string) (int, error) {
	r := new(model.CreatedResponse)
	r.Body.Status = http.StatusText(http.StatusCreated)
	r.Body.RecordID = recordID

	// Write the content.
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(r.Body)

	return http.StatusCreated, nil
}
