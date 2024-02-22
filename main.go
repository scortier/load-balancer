package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	listener, err := net.Listen("tcp", ":8080") // Change port to 80 for HTTP
	if err != nil {
		panic(err)
	}
	fmt.Println("Load Balancer is running on port 8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Failed to accept connection: %s\n", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read the request
	request, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		fmt.Printf("Failed to read request: %s\n", err)
		return
	}
	dump, err := httputil.DumpRequest(request, true)
	if err != nil {
		fmt.Printf("Failed to dump request: %s\n", err)
		return
	}
	fmt.Printf("%s\n", dump)

	// Forward the request to the backend server
	forwardRequest(conn, dump)
}

func forwardRequest(clientConn net.Conn, request []byte) {
	backendConn, err := net.Dial("tcp", "localhost:8081") // Backend server address
	if err != nil {
		fmt.Printf("Failed to connect to backend server: %s\n", err)
		return
	}
	defer backendConn.Close()

	// Send the request to the backend server
	backendConn.Write(request)

	// Read the response from the backend server
	io.Copy(clientConn, backendConn)
}
