package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := 8081 // Use a different port than the load balancer

	http.HandleFunc("/", handler)
	fmt.Printf("Backend Server is listening on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request from %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "HTTP/1.1 200 OK\n\nHello From Backend Server")
	fmt.Println("Replied with a hello message")
}
