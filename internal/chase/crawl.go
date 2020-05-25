package chase

import (
	"io"
)

//Chase fetches the relevant report urls per APT alias
func Chase(requestBody io.ReadCloser) {
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
	requestBody := TRAMRequest{"insert_report", "", ""}
	SendToTRAM(requestBody)
}
