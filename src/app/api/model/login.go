package model

// LoginResponse returns 200.
// swagger:response LoginResponse
type LoginResponse struct {
	// in: body
	Body struct {
		// Required: true
		Status string `json:"status"`
		// Required: true
		Token string `json:"token"`
	}
}
