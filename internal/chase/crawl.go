package chase

//Chase fetches the relevant report urls
func Chase(token string) {
	getReports(token)
}

func getReports(token string) {
	//TODO: Implement me

	/*
		We might be able to utilize the feedly API here and parse
		the results of the blogs...this is definitely easier
		than writing specific parsing code for each security blog.

		https://cloud.feedly.com/v3/streams/contents?streamId=feed%2Fhttp%3A%2F%2Fwww.fireeye.com%2Fblog%2Ffeed

		oauth: A2jGKuUNsybN2Wmsg7-lPcTxVdRJjpMb6at4mVAue3LMrKKiQpezqjDfJhpRFOAnmFaIOKq_7iSCra78-fUEmDFIOBOeHw6IJukFACG1IaWFWXnvBFEZprwpQ4NrZ1iiGzbO2Dvo75Dk4W2UA58OBndDydV8WNNWaw8sfaMry25j4hFYCv0O3MeGH9LFcU0kLWwWKR7sfsWkrf0InqQ3c50Gya0FnrW2DSHkJb1IVTmc12zMo6IacG4xqDC9:feedlydev

	*/

	/*
		This function shouldn't return anything, rather it
		will be invoked as a goroutine and call SendToTRAM.

		PSUEDOCODE:

		for each url in config.json {
			feedlyResp := go FetchBlogs(url)
			for each Blog in feedlyResp.Items {
				var request TRAMRequest
				request.title = feedlyResp.Items[i].Blogs.Title
				request.url = feedlyResp.Items[i].Blogs.OriginID
				go SendToTRAM(request)
			}

		}
	*/
	var feed int
	var blog int
	feeds := GetFeeds()
	for feed = range feeds {
		url := "http://cloud.freedly.com/v3/streams/contents?streamId=" + feeds[feed]
		feedlyResp := FetchBlogs(url, token)
		for blog = range feedlyResp.Items {
			var tramRequest TRAMRequest
			tramRequest.title = feedlyResp.Items[blog].Title
			tramRequest.title = feedlyResp.Items[blog].OriginID
			go SendToTRAM(tramRequest)
		}
	}
}
