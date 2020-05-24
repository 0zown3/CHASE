package chase

import (
	"io"
)

//Chase fetches the relevant report urls per APT alias
func Chase(requestBody io.ReadCloser) {
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
		go getReports(aliases[i])
	}
}

func getReports(alias string) {
	//TODO: Implement me

	/*
		This function shouldn't return anything, rather it
		will be invoked as a goroutine and call SendToTRAM.

		PSUEDOCODE:

		request := TRAMRequest{'insert_report', 'none', 'none'}

		crawling code here... (colly?)
		parse crawling response
		request.title = parsing result
		request.url = parsing result
		SendToTRAM(request)
	*/
	request := TRAMRequest{"insert_report", "", ""}
	SendToTRAM(request)
}
