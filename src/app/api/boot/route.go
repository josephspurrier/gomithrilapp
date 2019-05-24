package boot

import "app/api/endpoint"

// LoadRoutes will load the routes for the endpoints.
func LoadRoutes(core endpoint.Core) {
	endpoint.SetupStatic(core)
	endpoint.SetupLogin(core)
	endpoint.SetupRegister(core)
	endpoint.SetupNotepad(core)
}
