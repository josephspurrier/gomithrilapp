package endpoint

import (
	"errors"
	"net/http"

	"app/api/model"
	"app/api/pkg/structcopy"
)

// NoteEndpoint .
type NoteEndpoint struct {
	Core
}

// SetupNotepad .
func SetupNotepad(c Core) {
	p := new(NoteEndpoint)
	p.Core = c

	p.Router.Post("/v1/note", p.Create)
	p.Router.Get("/v1/note", p.Index)
	p.Router.Get("/v1/note/:note_id", p.Show)
	p.Router.Put("/v1/note/:note_id", p.Update)
	p.Router.Delete("/v1/note/:note_id", p.Destroy)
}

// Create .
// swagger:route POST /v1/note note NoteCreate
//
// Create a note.
//
// Security:
//   token:
//
// Responses:
//   201: CreatedResponse
//   400: BadRequestResponse
//   401: UnauthorizedResponse
//   500: InternalServerErrorResponse
func (p *NoteEndpoint) Create(w http.ResponseWriter, r *http.Request) (int, error) {
	// swagger:parameters NoteCreate
	type request struct {
		// in: body
		Body struct {
			Message string `json:"message"`
		}
	}

	// Request validation.
	req := new(request)
	if err := p.Bind.UnmarshalAndValidate(req, r); err != nil {
		return http.StatusBadRequest, err
	}

	// Get the user ID.
	userID, ok := p.Context.UserID(r)
	if !ok {
		return http.StatusInternalServerError, errors.New("invalid user")
	}

	// Create the note.
	ID, err := p.Store.Note.Create(userID, req.Body.Message)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return p.Response.Created(w, ID)
}

// Index .
// swagger:route GET /v1/note note NoteIndex
//
// List notes.
//
// Security:
//   token:
//
// Responses:
//   200: NoteIndexResponse
//   400: BadRequestResponse
//   401: UnauthorizedResponse
//   500: InternalServerErrorResponse
func (p NoteEndpoint) Index(w http.ResponseWriter, r *http.Request) (int, error) {
	// Get the user ID.
	userID, ok := p.Context.UserID(r)
	if !ok {
		return http.StatusInternalServerError, errors.New("invalid user")
	}

	// Get a list of notes for the user.
	group := p.Store.Note.NewGroup()
	_, err := p.Store.Note.FindAllByUser(&group, userID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Copy the items to the JSON model.
	arr := make([]model.Note, 0)
	for _, u := range group {
		item := new(model.Note)
		err = structcopy.ByTag(&u, "db", item, "json")
		if err != nil {
			return http.StatusInternalServerError, err
		}
		arr = append(arr, *item)
	}

	// Create the response.
	m := new(model.NoteIndexResponse).Body
	m.Notes = arr

	return p.Response.JSON(w, m)
}

// Show .
// swagger:route GET /v1/note/{note_id} note NoteShow
//
// Show a note.
//
// Security:
//   token:
//
// Responses:
//   200: NoteShowResponse
//   400: BadRequestResponse
//   401: UnauthorizedResponse
//   500: InternalServerErrorResponse
func (p NoteEndpoint) Show(w http.ResponseWriter, r *http.Request) (int, error) {
	// swagger:parameters NoteShow
	type request struct {
		// in: path
		NoteID string `json:"note_id" validate:"required"`
	}

	// Request validation.
	req := new(request)
	if err := p.Bind.UnmarshalAndValidate(req, r); err != nil {
		return http.StatusBadRequest, err
	}

	// Get the user ID.
	userID, ok := p.Context.UserID(r)
	if !ok {
		return http.StatusInternalServerError, errors.New("invalid user")
	}

	// Get the note for the user.
	note := p.Store.Note.New()
	exists, err := p.Store.Note.FindOneByIDAndUser(&note, req.NoteID, userID)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if !exists {
		return http.StatusBadRequest, errors.New("invalid note")
	}

	// Copy the items to the JSON model.
	item := new(model.Note)
	err = structcopy.ByTag(&note, "db", item, "json")
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Create the response.
	m := new(model.NoteShowResponse).Body
	m.Note = *item

	return p.Response.JSON(w, m)
}

// Update .
// swagger:route PUT /v1/note/{note_id} note NoteUpdate
//
// Update a note.
//
// Security:
//   token:
//
// Responses:
//   200: OKResponse
//   400: BadRequestResponse
//   401: UnauthorizedResponse
//   500: InternalServerErrorResponse
func (p *NoteEndpoint) Update(w http.ResponseWriter, r *http.Request) (int, error) {
	// swagger:parameters NoteUpdate
	type request struct {
		// in: path
		NoteID string `json:"note_id" validate:"required"`
		// in: body
		Body struct {
			Message string `json:"message"`
		}
	}

	// Request validation.
	req := new(request)
	if err := p.Bind.UnmarshalAndValidate(req, r); err != nil {
		return http.StatusBadRequest, err
	}

	// Get the user ID.
	userID, ok := p.Context.UserID(r)
	if !ok {
		return http.StatusInternalServerError, errors.New("invalid user")
	}

	// Determine if the note exists for the user.
	note := p.Store.Note.New()
	exists, err := p.Store.Note.FindOneByIDAndUser(&note, req.NoteID, userID)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if !exists {
		return http.StatusBadRequest, errors.New("note does not exist")
	}

	// Update the note.
	_, err = p.Store.Note.Update(req.NoteID, userID, req.Body.Message)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return p.Response.OK(w, "note updated")
}

// Destroy .
// swagger:route DELETE /v1/note/{note_id} note NoteDestroy
//
// Delete a note.
//
// Security:
//   token:
//
// Responses:
//   200: OKResponse
//   400: BadRequestResponse
//   401: UnauthorizedResponse
//   500: InternalServerErrorResponse
func (p NoteEndpoint) Destroy(w http.ResponseWriter, r *http.Request) (int, error) {
	// swagger:parameters NoteDestroy
	type request struct {
		// in: path
		NoteID string `json:"note_id" validate:"required"`
	}

	// Request validation.
	req := new(request)
	if err := p.Bind.UnmarshalAndValidate(req, r); err != nil {
		return http.StatusBadRequest, err
	}

	// Get the user ID.
	userID, ok := p.Context.UserID(r)
	if !ok {
		return http.StatusInternalServerError, errors.New("invalid user")
	}

	// Get a the note for the user.
	note := p.Store.Note.New()
	affected, err := p.Store.Note.DeleteOneByIDAndUser(&note, req.NoteID, userID)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if affected == 0 {
		return http.StatusBadRequest, errors.New("note does not exist")
	}

	return p.Response.OK(w, "note deleted")
}
