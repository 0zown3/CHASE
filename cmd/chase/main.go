package main

import (
	"chase/internal/chase"
	"fmt"
	"log"
	"net/http"
)

func main() {
	serverHandler := http.HandlerFunc(chase.Server)
	fmt.Print("Running CHASE on localhost:7000")
	if err := http.ListenAndServe(":7000", serverHandler); err != nil {
		log.Fatalf("Unable to start CHASE server on port 7000 %v", err)
	}
}
