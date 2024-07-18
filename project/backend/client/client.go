package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	//connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to the server", err)
		os.Exit(1)
	}
	defer conn.Close()

	// read input from the user and send it to the server
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error message: ")
			break
		}
		// send the message to the server
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message: ", err)
			break
		}
		// read the response from the server
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(" Error reading from server: ", err)
			break
		}

		fmt.Print("Response from server: ", response)
	}
}
