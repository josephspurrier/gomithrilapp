package component

// LoadRoutes will load the routes for the endpoints.
func LoadRoutes(core Core) {
	SetupStatic(core)
	SetupLogin(core)
	SetupRegister(core)
}
