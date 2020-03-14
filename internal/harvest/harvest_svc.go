//This is the "module" that will handle report gathering

package harvest

import (
	"log"
	"net/http"
)

//Report requires a url string
type Report struct {
	Url string
}

// Ping makes a HEAD request to the target URL
func Ping(reportURL string) int {
	resp, err := http.Head(reportURL)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode == 200 {
		return resp.StatusCode
	}
	return 0
}
