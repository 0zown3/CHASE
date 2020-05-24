package chase

//SendToTRAM communicates with the TRAM endpoint that ingests report urls
func SendToTRAM(tramRequest TRAMRequest) {
	//TODO: Implement me
	/*
		We need to POST to the endpoint defined here:
		https://github.com/mitre-attack/tram/blob/master/handlers/web_api.py#L30

		this would be /rest

		the /rest endpoint is looking for:

		data = {
			'index': 'insert_report',
			'title': str
			'url': str
		}

		within insert_report, title and url are used.

	*/
}
