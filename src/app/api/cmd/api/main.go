package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"app/api/boot"
	"app/api/component"
	"app/api/middleware"
	"app/api/pkg/logger"
)

func init() {
	// Verbose logging with file name and line number.
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores.
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	port := "8081"

	// Create the logger.
	//l := logger.New(log.New(os.Stderr, "", log.LstdFlags))
	l := logger.New(log.New(os.Stderr, "", log.Lshortfile))

	// Setup the services.
	core := boot.Services(l)

	// Load the routes.
	boot.SetupRouter(l, core.Router)
	component.LoadRoutes(core)

	// Start the web server.
	l.Printf("Server started.")
	err := http.ListenAndServe(":"+port, middleware.Wrap(core.Router, l, core.Token.Secret()))
	if err != nil {
		l.Printf(err.Error())
	}
}
