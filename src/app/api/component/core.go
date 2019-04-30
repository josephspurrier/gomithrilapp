package component

import "github.com/husobee/vestigo"

// Core represents the core for all the endpoints.
type Core struct {
	Router *vestigo.Router
}

// NewCore returs a core for all the endpoints.
func NewCore() Core {
	return Core{
		Router: vestigo.NewRouter(),
	}
}
