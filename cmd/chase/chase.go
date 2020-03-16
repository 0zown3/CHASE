//CHASE
//Concurrent Harvester for Adversary Studies & Examination

package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", StartHarvest)
	log.Print("CHASE running @ localhost:7000")
	log.Fatal(http.ListenAndServe(":7000", nil))
}
