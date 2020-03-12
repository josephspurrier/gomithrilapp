package main

import (
	"log"
	"net/http"
	"os"

	"app/api/config"
	"app/api/pkg/logger"
	"app/api/pkg/mock"
)

func main() {
	// FIXME: This should be an environment variable.
	port := "8081"

	// Create the logger.
	//l := logger.New(log.New(os.Stderr, "", log.LstdFlags))
	l := logger.New(log.New(os.Stderr, "", log.Lshortfile))

	// Setup the services.
	core := config.Services(l, config.Database(l), mock.New(false))
	config.LoadRoutes(core)

	// Start the web server.
	l.Printf("Server started.")
	err := http.ListenAndServe(":"+port, config.Middleware(core))
	if err != nil {
		l.Printf(err.Error())
	}
}
