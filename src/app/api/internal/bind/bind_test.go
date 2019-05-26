package bind_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"app/api/internal/bind"
	"app/api/internal/testrequest"
	"app/api/pkg/mock"
	"app/api/pkg/router"

	"github.com/stretchr/testify/assert"
)

func TestFormSuccess(t *testing.T) {
	called := false

	mux := router.New()

	mux.Post("/user/:user_id", router.Handler(
		func(w http.ResponseWriter, r *http.Request) (status int, err error) {
			called = true

			// swagger:parameters UserCreate
			type request struct {
				// in: path
				UserID string `json:"user_id" validate:"required"`
				// in: formData
				// Required: true
				FirstName string `json:"first_name" validate:"required"`
				// in: formData
				// Required: true
				LastName string `json:"last_name" validate:"required"`
			}

			req := new(request)
			b := bind.New(mock.New(false), mux)

			assert.Nil(t, b.Unmarshal(&req, r))
			assert.Nil(t, b.Validate(req))

			assert.Equal(t, "10", req.UserID)
			assert.Equal(t, "john", req.FirstName)
			assert.Equal(t, "smith", req.LastName)
			return http.StatusOK, nil
		}))

	form := url.Values{}
	form.Add("first_name", "john")
	form.Add("last_name", "smith")

	r := httptest.NewRequest("POST", "/user/10", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, true, called)
}

func TestFormNil(t *testing.T) {
	called := false

	mux := router.New()

	mux.Post("/user/:user_id", router.Handler(
		func(w http.ResponseWriter, r *http.Request) (status int, err error) {
			called = true

			// swagger:parameters UserCreate
			type request struct {
				// in: path
				UserID string `json:"user_id" validate:"required"`
				// in: formData
				// Required: true
				FirstName string `json:"first_name" validate:"required"`
				// in: formData
				// Required: true
				LastName string `json:"last_name" validate:"required"`
			}

			req := request{}
			b := bind.New(mock.New(false), mux)

			assert.NotNil(t, b.Unmarshal(req, r))

			assert.Equal(t, "", req.UserID)
			assert.Equal(t, "", req.FirstName)
			assert.Equal(t, "", req.LastName)
			return http.StatusOK, nil
		}))

	r := httptest.NewRequest("POST", "/user/10", nil)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, true, called)
}

func TestFormMissingPointer(t *testing.T) {
	called := false

	mux := router.New()

	mux.Post("/user/:user_id", router.Handler(
		func(w http.ResponseWriter, r *http.Request) (status int, err error) {
			called = true

			// swagger:parameters UserCreate
			type request struct {
				// in: path
				UserID string `json:"user_id" validate:"required"`
				// in: formData
				// Required: true
				FirstName string `json:"first_name" validate:"required"`
				// in: formData
				// Required: true
				LastName string `json:"last_name" validate:"required"`
			}

			req := request{}
			b := bind.New(mock.New(false), mux)

			assert.NotNil(t, b.Unmarshal(req, r))

			assert.Equal(t, "", req.UserID)
			assert.Equal(t, "", req.FirstName)
			assert.Equal(t, "", req.LastName)
			return http.StatusOK, nil
		}))

	form := url.Values{}
	form.Add("first_name", "john")
	form.Add("last_name", "smith")

	r := httptest.NewRequest("POST", "/user/10", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, true, called)
}

func TestJSONSuccess(t *testing.T) {
	called := false

	mux := router.New()

	mux.Post("/user/:user_id", router.Handler(
		func(w http.ResponseWriter, r *http.Request) (status int, err error) {
			called = true

			// swagger:parameters UserCreate
			type request struct {
				// in: body
				Body struct {
					// in: path
					UserID string `json:"user_id" validate:"required"`
					// Required: true
					FirstName string `json:"first_name" validate:"required"`
					// Required: true
					LastName string `json:"last_name" validate:"required"`
				}
			}

			req := new(request).Body

			b := bind.New(mock.New(false), mux)

			assert.Nil(t, b.Unmarshal(&req, r))
			assert.Nil(t, b.Validate(req))

			assert.Equal(t, "10", req.UserID)
			assert.Equal(t, "john", req.FirstName)
			assert.Equal(t, "smith", req.LastName)
			return http.StatusOK, nil
		}))

	form := url.Values{}
	form.Add("first_name", "john")
	form.Add("last_name", "smith")

	r := httptest.NewRequest("POST", "/user/10", strings.NewReader(testrequest.ToJSON(form)))
	r.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, true, called)
}

func TestJSONFailure(t *testing.T) {
	called := false

	mux := router.New()

	mux.Post("/user/:user_id", router.Handler(
		func(w http.ResponseWriter, r *http.Request) (status int, err error) {
			called = true

			// swagger:parameters UserCreate
			type request struct {
				// in: body
				Body struct {
					// in: path
					UserID string `json:"user_id" validate:"required"`
					// Required: true
					FirstName string `json:"first_name" validate:"required"`
					// Required: true
					LastName string `json:"last_name" validate:"required"`
				}
			}

			req := new(request).Body

			b := bind.New(mock.New(false), mux)

			assert.NotNil(t, b.Unmarshal(req, r))

			assert.Equal(t, "", req.UserID)
			assert.Equal(t, "", req.FirstName)
			assert.Equal(t, "", req.LastName)
			return http.StatusOK, nil
		}))

	form := url.Values{}
	form.Add("first_name", "john")
	form.Add("last_name", "smith")

	r := httptest.NewRequest("POST", "/user/10", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, true, called)
}

func TestJSONFailureNil(t *testing.T) {
	called := false

	mux := router.New()

	mux.Post("/user/:user_id", router.Handler(
		func(w http.ResponseWriter, r *http.Request) (status int, err error) {
			called = true

			// swagger:parameters UserCreate
			type request struct {
				// in: body
				Body struct {
					// in: path
					UserID string `json:"user_id" validate:"required"`
					// Required: true
					FirstName string `json:"first_name" validate:"required"`
					// Required: true
					LastName string `json:"last_name" validate:"required"`
				}
			}

			req := new(request).Body

			b := bind.New(mock.New(false), mux)

			assert.NotNil(t, b.Unmarshal(req, r))

			assert.Equal(t, "", req.UserID)
			assert.Equal(t, "", req.FirstName)
			assert.Equal(t, "", req.LastName)
			return http.StatusOK, nil
		}))

	r := httptest.NewRequest("POST", "/user/10", nil)
	r.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, true, called)
}

func TestMock(t *testing.T) {
	b := bind.New(mock.New(true), nil)

	e := errors.New("test error")

	b.Mock.Add("Binder.Unmarshal", e)
	assert.Equal(t, e, b.Unmarshal(nil, nil))

	b.Mock.Add("Binder.Validate", e)
	assert.Equal(t, e, b.Validate(nil))
}
