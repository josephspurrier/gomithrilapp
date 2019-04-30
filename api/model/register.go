package model

// RegisterResponse returns 200.
// swagger:response RegisterResponse
type RegisterResponse struct {
	// in: body
	Body struct {
		// Required: true
		Status string `json:"status"`
		// Required: true
		Success bool `json:"success"`
	}
}
