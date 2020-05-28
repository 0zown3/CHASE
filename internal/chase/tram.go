package chase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//SendToTRAM communicates with the TRAM endpoint that ingests report urls
func SendToTRAM(tramBody TRAMRequest) {
	//endpoint can't remain hardcoded, needs to be more dynamic if deployed in container env
	endpoint := "http://localhost:9999/rest"
	contentType := "application/json"
	requestBody, _ := json.Marshal(tramBody)
	fmt.Printf("Sending %s report to TRAM!\n", tramBody.Title[0])
	http.Post(endpoint, contentType, bytes.NewBuffer(requestBody))
}
