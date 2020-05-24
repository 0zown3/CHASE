package chase

import (
	"io"
)

//GetReportUrls fetches the relevant report urls per APT alias
func GetReportUrls(requestBody io.ReadCloser) {
	/*
		We need to setup a channel in this function to handle
		responses appending to the urls slice. Additionally, in
		the for loop, we need to make a goroutine call to the
		function that fetches report urls for each alias.

		i.e.

		urls = append(urls, go crawl(aliases[i]))

		something along those lines should work..
	*/
	var urls []string
	apt := DecodeBody(requestBody)
	aliases := GetAliases(apt)
	for i := range aliases {
		urls = append(urls, aliases[i])
	}
	SendToTRAM(urls)
}
