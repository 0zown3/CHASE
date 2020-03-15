//This is the "module" that will handle report gathering

package harvest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//Report requires a url string
type Report struct {
	URL string
}

//Domains is the struct used for unmarshalling the allowed_domains json file.
type Domains struct {
	Domains []string `json:"domains"`
}

func loadDomains() []string {
	var domains Domains
	jsonFile, _ := os.Open("../../config/allowed_domains.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &domains)
	return domains.Domains
}

//FetchReports accepts a dtg and apt string
func FetchReports(dtg string, apt string) {
	allowedDomains := loadDomains()
	fmt.Println(allowedDomains)
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
