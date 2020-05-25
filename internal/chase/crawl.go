package chase

//Chase fetches the relevant report urls
func Chase() {
	getReports()
}

func getReports() {
	//TODO: Implement me

	//https://www.google.com/search?q=Fancy+Bear&num=10000
	//painfully parse these results..?

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
