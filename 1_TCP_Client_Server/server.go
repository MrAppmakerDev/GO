// 2. server.go - Server Implementation
// This file contains all server-related functionality:

// Setting up a TCP listener
// Accepting connections
// Reading and displaying incoming messages
// Handling connection errors

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

// RunServer starts a TCP server that listens for incoming messages
func RunServer(address string) {
	fmt.Println("Server listening on", address)
	fmt.Println("Waiting for client to connect...")

	// Create TCP listener
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error listening: ", err)
		os.Exit(1)
	}
	defer listener.Close()

	// Accept connection
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err)
		os.Exit(1)
	}
	fmt.Println("Client connected: ", conn.RemoteAddr())

	// Set up reader for incoming messages
	reader := bufio.NewReader(conn)

	// Handle incoming messages
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed or error reading:", err)
			break
		}

		timestamp := time.Now().Format("15:04:05")
		fmt.Printf("[%s] %s", timestamp, message)
	}
}