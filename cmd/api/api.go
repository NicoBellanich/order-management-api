// Package main is the entry point for the microblogging platform API server.
// It initializes and starts the HTTP server.
package main

import (
	"github.com/nicobellanich/order-management-api/cmd/server"
)

// main starts the API server.
func main() {
	server.Run()
}
