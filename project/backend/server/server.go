package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("error starting server :", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("server is listening on port 8080")

	for {
		// accpet an incoming connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		// handle the connection in a new goroutine
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("New connection accepted")

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			break
		}
		// print the received message
		fmt.Print("message received:", message)

		//echo the message back to the client
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			break
		}
	}
}
