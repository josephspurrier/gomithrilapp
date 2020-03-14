// Package main is a reverse proxy that serves both the UI and API depending
// on the URL path.
package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

const (
	defaultUI  = "http://ui:8080"
	defaultAPI = "http://api:8081"
)

func main() {
	p := NewProxy(defaultUI, defaultAPI)
	http.HandleFunc("/", p.handle)
	log.Fatalln(http.ListenAndServe(":80", nil))
}

// Proxy is a proxy for the UI and API.
type Proxy struct {
	targetUI  *url.URL
	targetAPI *url.URL
	proxyUI   *httputil.ReverseProxy
	proxyAPI  *httputil.ReverseProxy
}

// NewProxy returns a new instance of the proxy.
func NewProxy(targetUI, targetAPI string) *Proxy {
	// Get the URLs from environment variables if they are set.
	envUI := os.Getenv("UI_URL")
	envAPI := os.Getenv("API_URL")
	if len(envUI) > 0 {
		targetUI = envUI
	}
	if len(envAPI) > 0 {
		targetAPI = envAPI
	}

	// Convert the strings into URLs.
	ui, err := url.Parse(targetUI)
	if err != nil {
		log.Println("UI Error:", err)
	}
	api, err := url.Parse(targetAPI)
	if err != nil {
		log.Println("API Error:", err)
	}

	return &Proxy{
		targetUI:  ui,
		targetAPI: api,
		proxyUI:   httputil.NewSingleHostReverseProxy(ui),
		proxyAPI:  httputil.NewSingleHostReverseProxy(api),
	}
}

func (p *Proxy) handle(w http.ResponseWriter, r *http.Request) {
	// If the path starts with /api, serve the API.
	if strings.HasPrefix(r.URL.Path, "/api") {
		p.proxyAPI.ServeHTTP(w, r)
		return
	}

	// Else, serve the UI.
	p.proxyUI.ServeHTTP(w, r)
}
