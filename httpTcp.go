package main

import (
	"fmt"
	"net"
	"net/http"
)

func httpTcp() {

}

// TCP
func TcpConnection() {
	// 1. Start listening on port 8080
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("TCP Server running on port 8080...")

	for {
		// 2. Accept connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err)
			continue
		}

		// 3. Handle connection in goroutine
		handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Println("Received:", string(buffer[:n]))

	// 5. Write response
	conn.Write([]byte("Hello from TCP Server!"))
}

// | Concept             | Meaning                              |
// | ------------------- | ------------------------------------ |
// | `net.Listen()`      | Start server                         |
// | `listener.Accept()` | Accept client                        |
// | `net.Conn`          | Connection object                    |
// | `Read()`            | Receive data                         |
// | `Write()`           | Send data                            |
// | `goroutine`         | Handle multiple clients concurrently |

// HTTP server is built on top of TCP.
// Define routes (URLs), Write handler functions, Start server

// HTTP
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from HTTP Server!")
}

func hhtp() {
	// Register route
	http.HandleFunc("/", homeHandler)

	fmt.Println("HTTP Server running on port 8080...")

	// Start server
	http.ListenAndServe(":8080", nil)
}

// | Component               | Purpose          |
// | ----------------------- | ---------------- |
// | `http.HandleFunc()`     | Register route   |
// | `http.ResponseWriter`   | Send response    |
// | `*http.Request`         | Get request data |
// | `http.ListenAndServe()` | Start server     |

// Working: Opens TCP connection, Parses HTTP request, Calls matching handler, Writes HTTP response
