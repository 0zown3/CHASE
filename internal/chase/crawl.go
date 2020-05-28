package chase

//Chase fetches the relevant report urls
func Chase(token string) {
	getReports(token)
}

func getReports(token string) {
	var feed int
	feeds := GetFeeds()
	for feed = range feeds {
		url := "https://cloud.feedly.com/v3/streams/contents?streamId=" + feeds[feed]
		feedlyResp := FetchBlogs(url, token)
		go feedTram(feedlyResp.Items)
	}
}

func feedTram(blogs []Blogs) {
	var blog int
	for blog = range blogs {
		var tramRequest TRAMRequest
		tramRequest.Index = "insert_report"
		tramRequest.Title[0] = blogs[blog].Title
		tramRequest.URL[0] = blogs[blog].OriginID
		go SendToTRAM(tramRequest)
	}
}
