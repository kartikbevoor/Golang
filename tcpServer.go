package main

import (
	"fmt"
	"net"
)

// TCP (Transmission Control Protocol) is: Connection-oriented, Reliable (guarantees delivery)
// Ordered data transmission, Error-checked

// Working
// 1. Listen on a port
// 2. Accept client connections
// 3. Read/Write data
// 4. Close connection
func tcpServer() {

}

// basic tcp server
func basicTcpServer() {
	// 1. Start listening on port 8080
	listener, err := net.Listen("tcp", ":8080") // net.Listen(network, address)
	// "tcp" → We are creating a TCP server. ":8080" → Listen on port 8080 on all available interfaces.
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP Server started on port 8080...")

	for {
		// 2. Accept client connection
		conn, err := listener.Accept() // Accept() blocks (waits) until a client connects.
		// Returns: conn → the connection object, err → error if something fails
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("Client connected:", conn.RemoteAddr()) // RemoteAddr() gives client IP and port.

		// 3. Handle client in goroutine
		go handleConnectionHttpClient(conn) // Handle Each Client Concurrently
	}
}

func handleConnectionHttpClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	// 4. Read data
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Println("Received:", string(buffer[:n]))

	// 5. Send response
	conn.Write([]byte("Message received!"))
}

// TCP Server Architecture
// +-------------------+
// Client 1  --->  |                   |
// Client 2  --->  |   TCP SERVER      |
// Client 3  --->  |                   |
//                 +-------------------+
//                         |
//                  Goroutine per client

// Handling Multiple Messages (Continuous Read)
// for {
// 	n, err := conn.Read(buffer)
// 	if err != nil {
// 		return
// 	}
// 	fmt.Println(string(buffer[:n]))
// }
