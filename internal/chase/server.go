package chase

import (
	"fmt"
	"net/http"
)

//Server is the handler function for the CHASE server
//that returns a slice of all APT related urls
func Server(writer http.ResponseWriter, request *http.Request) {
	apt := DecodeBody(request.Body)
	aliases := GetAliases(apt)
	/*
		The general flow of operations at this point, after retrieving
		the APT parsed from the request, is that we first retrieve
		the array of APT aliases from the apt_mappings.json file.

		These will be the aliases we use to search for the corresponding
		reports on the Internet. For example, if we query CHASE with APT28
		and only search on APT28, we will miss reports discussing Fancy Bear,
		which is the same group.

		After we retrieve this array, we concurrently spider the Internet
		for relevant URLs and return those aggregated results.

		Since we will be accessing the same slice resource to append
		URLs, we need to create channels to avoid deadlocks.

		PSEUDOCODE:

		slice = []
		aliases = getAliases(apt) <<< returns an array of strings
		for alias in aliases:
			go getReportUrls(alias)

		return slice
	*/

	fmt.Fprint(writer, apt)
}
