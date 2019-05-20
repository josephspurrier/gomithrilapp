package main

import (
	"log"
	"net/http"
	"os"

	"app/api/boot"
	"app/api/middleware"
	"app/api/pkg/logger"
)

func main() {
	// FIXME: This should be an environment variable.
	port := "8081"

	// Create the logger.
	//l := logger.New(log.New(os.Stderr, "", log.LstdFlags))
	l := logger.New(log.New(os.Stderr, "", log.Lshortfile))

	// Setup the services.
	core := boot.Services(l)

	// Start the web server.
	l.Printf("Server started.")
	err := http.ListenAndServe(":"+port, middleware.Wrap(core.Router, l, core.Token.Secret()))
	if err != nil {
		l.Printf(err.Error())
	}
}
