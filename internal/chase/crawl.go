package chase

//Chase fetches the relevant report urls
func Chase(token string) {
	getReports(token)
}

func getReports(token string) {
	var feed int
	var blog int
	feeds := GetFeeds()
	for feed = range feeds {
		url := "http://cloud.feedly.com/v3/streams/contents?streamId=" + feeds[feed]
		feedlyResp := FetchBlogs(url, token)
		for blog = range feedlyResp.Items {
			var tramRequest TRAMRequest
			tramRequest.title = feedlyResp.Items[blog].Title
			tramRequest.url = feedlyResp.Items[blog].OriginID
			go SendToTRAM(tramRequest)
		}
	}
}
