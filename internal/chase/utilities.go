package chase

import (
	"encoding/json"
	"io"
	"net/http"
)

//DecodeBody decodes the request body
func DecodeBody(requestBody io.ReadCloser) string {
	var body Request
	json.NewDecoder(requestBody).Decode(&body)
	return body.Token
}

//FetchBlogs is a generic utility that fetches blog posts from feedly
func FetchBlogs(url string) FeedlyResponse {
	var feedlyResp FeedlyResponse
	response, _ := http.Get(url)
	json.NewDecoder(response.Body).Decode(&feedlyResp)
	return feedlyResp
}
