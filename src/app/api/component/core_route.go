package component

import (
	"encoding/json"
	"net/http"

	"app/api/model"
	"app/api/pkg/router"
)

// LoadRoutes will load the endpoints.
func LoadRoutes(core Core) {
	SetupStatic(core)
	SetupLogin(core)
	SetupRegister(core)

	// Set up the 404 page.
	core.Router.Instance().NotFound = router.Handler(
		func(w http.ResponseWriter, r *http.Request) (int, error) {
			return http.StatusNotFound, nil
		})

	// Set the handling of all responses.
	router.ServeHTTP = func(w http.ResponseWriter, r *http.Request, status int, err error) {
		// Handle only errors.
		if status >= 400 {
			resp := new(model.GenericResponse)
			resp.Body.Status = http.StatusText(status)
			if err != nil {
				resp.Body.Message = err.Error()
			}

			// Write the content.
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(status)
			err := json.NewEncoder(w).Encode(resp.Body)
			if err != nil {
				w.Write([]byte(`{"status":"Internal Server Error","message":"problem encoding JSON"}`))
				return
			}
		}

		// Display server errors.
		if status >= 500 {
			if err != nil {
				core.Log.Printf("%v", err)
			}
		}
	}
}
