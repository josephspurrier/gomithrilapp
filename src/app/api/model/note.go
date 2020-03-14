package model

// NoteIndexResponse returns 200.
// swagger:response NoteIndexResponse
type NoteIndexResponse struct {
	// in: body
	Body struct {
		// Required: true
		Notes []Note `json:"notes"`
	}
}

// Note is a note of a user.
type Note struct {
	// Required: true
	UserID string `json:"id"`
	// Required: true
	Message string `json:"message"`
}

// NoteShowResponse returns 200.
// swagger:response NoteShowResponse
type NoteShowResponse struct {
	// in: body
	Body struct {
		// Required: true
		Note
	}
}
