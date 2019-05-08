package main

import (
	"./src/httpserver"
	"log"
)

const port = 8080

func main() {
	log.Printf("Starting server on port %d\n", port)

	httpserver.Start(port)
}
