package component

// LoadRoutes will load the endpoints.
func LoadRoutes(core Core) {
	SetupStatic(core)
	SetupLogin(core)
	SetupRegister(core)
}
