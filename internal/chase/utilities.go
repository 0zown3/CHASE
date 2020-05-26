package chase

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/oauth2"
)

//GetFeeds returns supported Feedly research feeds
func GetFeeds() []string {
	var feeds Feeds
	jsonFile, _ := os.Open("../../config/feeds.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &feeds)
	return feeds.Feeds
}

//DecodeBody decodes the request body
func DecodeBody(requestBody io.ReadCloser) string {
	var body Request
	json.NewDecoder(requestBody).Decode(&body)
	return body.Token
}

//FetchBlogs is a generic utility that fetches blog posts from feedly
func FetchBlogs(url string, token string) FeedlyResponse {
	var feedlyResp FeedlyResponse
	ctx := context.Background()
	client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: token,
		TokenType:   "Bearer",
	}))
	response, _ := client.Get(url)
	json.NewDecoder(response.Body).Decode(&feedlyResp)
	return feedlyResp
}
