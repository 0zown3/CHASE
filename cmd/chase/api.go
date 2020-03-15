package main

import (
	"chase/internal/harvest"
	"encoding/json"
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
		harvest.FetchReports(dtg, apt)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"harvest_status": "invalid method"}`))
	}
}
