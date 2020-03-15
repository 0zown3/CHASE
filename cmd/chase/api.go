package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type data struct {
	DTG string
	APT string
}

//StartHarvest handles the POST request that kicks off a harvest.
func StartHarvest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"harvest_status": "started"`))
		decoder := json.NewDecoder(r.Body)
		var data data
		err := decoder.Decode(&data)
		if err != nil {
			panic(err)
		}
		dtg := data.DTG
		apt := data.APT
		fetchReports(dtg, apt)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"harvest_status": "invalid method"}`))
	}
}

func fetchReports(dtg string, apt string) {
	fmt.Println(dtg)
	fmt.Println(apt)
}
