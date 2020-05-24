package chase

import (
	"io"
)

//GetReportUrls fetches the relevant report urls per APT alias
func GetReportUrls(requestBody io.ReadCloser) {
	/*
		TRAM's /rest endpoint isn't expecting an array of urls to be passed
		to insert_reports. We need to send individual POST requests containing
		the title and url of each report.

		Therefore, in addition to creating goroutines for fetching reports per alias,
		we also need to concurrently send individual requests to the /rest endpoint
		for each URL we collect.

		So we can't pass a slice to SendToTRAM, we must define a new interface to
		store this data.
	*/
	apt := DecodeBody(requestBody)
	aliases := GetAliases(apt)
	for i := range aliases {
		//go gather(aliases[i])
		//something like that
	}
}
