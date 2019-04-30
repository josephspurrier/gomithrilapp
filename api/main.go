package main

import (
	"log"
	"net/http"

	"github.com/husobee/vestigo"
	"github.com/josephspurrier/govueapp/api/component"
	"github.com/josephspurrier/govueapp/api/middleware"
	"github.com/josephspurrier/govueapp/api/pkg/app"
)

func init() {
	// Verbose logging with file name and line number.
	log.SetFlags(log.Lshortfile)
}

func main() {
	port := "8081"

	router := LoadRoutes()
	router.SetGlobalCors(&vestigo.CorsAccessControl{
		AllowOrigin:  []string{"*"},
		AllowHeaders: []string{"Content-Type", "Origin", "X-Requested-With", "Accept"},
	})

	log.Println("Server started.")
	err := http.ListenAndServe(":"+port, middleware.Log(router))
	if err != nil {
		log.Println(err)
	}
}

// LoadRoutes will load the endpoints.
func LoadRoutes() *vestigo.Router {
	core := app.NewCore()

	component.SetupStatic(core)
	component.SetupLogin(core)
	component.SetupRegister(core)

	return core.Router
}
