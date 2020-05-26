package tests

import (
	"bytes"
	"chase/internal/chase"
	"encoding/json"
	"net/http"
	"testing"
)

//TestDecodeBody tests the DecodeBody utility function in internal/chase
func TestDecodeBody(t *testing.T) {
	t.Run("returns string from DecodeBody utility", func(t *testing.T) {
		requestBody := constructBody()
		jsonBody, _ := json.Marshal(requestBody)
		request, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))
		got := chase.DecodeBody(request.Body)
		want := "test123"
		assertEquals(t, got, want)
	})
}

//TestFetchBlogs mocks an API call to the Feedly API and tests the unmarshalled response
func TestFetchBlogs(t *testing.T) {
	t.Run("tests the function that makes a request to Feedly's API", func(t *testing.T) {
		var feedlyResp chase.FeedlyResponse
		testServer := feedlyResponseStub()
		defer testServer.Close()
		response, _ := http.Get(testServer.URL)
		json.NewDecoder(response.Body).Decode(&feedlyResp)
		assertEquals(t, feedlyResp.FeedTitle, "Test Title")
	})
}
