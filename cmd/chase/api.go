package main

import (
	"net/http"
)

//Harvest handles the POST request that kicks off a harvest.
func Harvest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "harvest started"`))
	}
}
