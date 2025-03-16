// 1. main.go - The Entry Point
// This file serves as the entry point of our application. It:

// Parses command-line arguments
// Validates input
// Delegates to the appropriate module (server or client)
// Defines constants used across the application

// ----------    Help    ----------

// How to Build and Run

// Place all three files in the same directory

// Go offers several ways to build your program:

// Option 1: Direct execution without explicit build
// go run *.go
// This command compiles and runs all Go files in the current directory in one step. 
// It's useful during development but creates a temporary executable.

// Option 2: Build a named executable
// go build -o echo-app
// This creates an executable named echo-app in your current directory. 
// The -o flag specifies the output filename.

// Step 4: Run the application
// After building with Option 2, you can run your program:
// For the server:
// ./echo-app -server -port 8080

// For the client (in a separate terminal window):
// ./echo-app -client -host localhost -port 8080

package main

import (
	"flag"
	"fmt"
	"os"
)

// Package-level constants that can be accessed by imported files
const (
	defaultPort = "8080"
	defaultHost = "localhost"
)

func main() {
	// Define command line flags
	serverMode := flag.Bool("server", false, "Run in server mode")
	clientMode := flag.Bool("client", false, "Run in client mode")
	port := flag.String("port", defaultPort, "Port to use for communication")
	host := flag.String("host", defaultHost, "Host address (for client machine)")
	flag.Parse() // Process the actual command-line arguments.

	// Validate arguments
	if *serverMode == *clientMode {
		fmt.Println("Error: You must specify either -server or -client mode")
		flag.Usage()
		os.Exit(1)
	}

	// Full address for connection
	address := fmt.Sprintf("%s: %s", *host, *port)

	// Run in appropriate mode
	if *serverMode {
		RunServer(address)
	} else {
		RunClient(address)
	}
}
