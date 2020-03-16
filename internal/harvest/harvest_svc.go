//This is the "module" that will handle report gathering

package harvest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

//Reports requires an array of strings
type Reports struct {
	URLS []string
}

//Domains is the struct used for unmarshalling the allowed_domains json file.
type Domains struct {
	Domains []string `json:"domains"`
}

//FetchReports accepts a dtg and apt string
func FetchReports(dtg string, apt string) {
	allowedDomains := loadDomains()
	for i := 0; i <= len(allowedDomains)-1; i++ {
		c := colly.NewCollector(
			colly.AllowedDomains(allowedDomains[i]),
			colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36"),
		)
		if allowedDomains[i] == "www.fireeye.com" {
			url := constructFireeyeURL(dtg, allowedDomains[i])
			fireeyeHarvester(url, apt, c)
		}
	}
}

//loadDomains loads allowed domains for crawling into memory
func loadDomains() []string {
	var domains Domains
	jsonFile, _ := os.Open("../../config/allowed_domains.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &domains)
	return domains.Domains
}

func constructFireeyeURL(dtg string, domain string) string {
	suffix := getFireeyeSuffix(dtg)
	prefix := "https://" + domain + "/blog/threat-research"
	return prefix + suffix
}

func getFireeyeSuffix(dtg string) string {
	year := dtg[:4]
	month := dtg[4:6]
	suffix := "/" + year + "/" + month + ".html"
	return suffix
}

func fireeyeHarvester(url string, apt string, c *colly.Collector) {
	c.OnHTML("h4", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, apt) {
			articleTitle := e.Text
			fmt.Println(articleTitle)
		}
	})
	log.Printf("Crawling %s for reports relating to %s", url, apt)
	c.Visit(url)
}
