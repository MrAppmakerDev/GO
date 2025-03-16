// 3. client.go - Client Implementation
// This file contains all client-related functionality:

// Connecting to a server
// Reading user input
// Sending messages to the server
// Handling disconnections

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// RunClient connects to a TCP server and sends user input
func RunClient(address string) {
	fmt.Println("Connecting to server at", address)

	// Connect to server
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected! Type messages and press Enter to send.")
	fmt.Println("Type 'exit' to quit.")

	// Read user input and send to server
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()

		if strings.ToLower(message) == "exit" {
			fmt.Println("Disconnecting...")
			break
		}

		// Send message to server
		_, err := fmt.Fprintln(conn, message)
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}
	}
}
