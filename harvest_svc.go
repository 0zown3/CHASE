//This is the "module" that will handle report gathering

package main

import (
	"log"
	"net/http"
)

//Report requires a url string
type Report struct {
	url string
}

func ping(reportURL string) int {
	resp, err := http.Head(reportURL)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode == 200 {
		return resp.StatusCode
	}
	return 0
}
