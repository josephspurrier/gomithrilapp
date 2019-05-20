package boot

import "app/api/component"

// LoadRoutes will load the routes for the endpoints.
func LoadRoutes(core component.Core) {
	component.SetupStatic(core)
	component.SetupLogin(core)
	component.SetupRegister(core)
}
