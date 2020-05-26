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

//TestFetchBlogs mocks an API call to the Feedly API and tests the unmarshalled
func TestFetchBlogs(t *testing.T) {
	t.Run("tests the function that makes a request to Feedly's API", func(t *testing.T) {
		testServer := feedlyResponseStub()
		defer testServer.Close()
		feedlyResp := chase.FetchBlogs(testServer.URL)
		assertEquals(t, feedlyResp.FeedTitle, "Test Title")
	})
}
