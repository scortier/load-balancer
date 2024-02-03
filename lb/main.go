package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	port := 8080 // You can change this to any port you prefer

	http.HandleFunc("/", handler)
	fmt.Printf("Load Balancer is listening on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request from %s\n", r.RemoteAddr)

	// Choose a backend server (in this case, just one backend server for simplicity)
	backendURL, _ := url.Parse("http://localhost:8081")
	proxy := httputil.NewSingleHostReverseProxy(backendURL)

	// Forward the request to the backend server
	proxy.ServeHTTP(w, r)
}
