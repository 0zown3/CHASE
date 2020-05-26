package tests

import (
	"bytes"
	"chase/internal/chase"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

/*
	This API provides a domain specific testing language for CHASE

	Below are the assertion functions used to verify responses and
	general utility functions that help keep tests clean.
*/

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Incorrect status: got %d, want %d", got, want)
	}
}

//TODO: CHANGE MY RETURN TYPE
func assertResponse(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Incorrect status: got %s, want %s", got, want)
	}
}

func assertEquals(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Incorrect status: got %s, want %s", got, want)
	}
}

func assertType(t *testing.T, got, want reflect.Type) {
	t.Helper()
	if got != want {
		t.Errorf("Incorrect types: got %T, want %T", got, want)
	}
}

/*
	Testing Utilities
*/

func constructBody() chase.Request {
	var request chase.Request
	request.Token = "test123"
	return request
}

func getRequestBody() io.ReadCloser {
	requestBody := constructBody()
	jsonBody, _ := json.Marshal(requestBody)
	request, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))
	return request.Body
}

//Creates mock api server to test feedly response
func feedlyResponseStub() *httptest.Server {
	var testResponse chase.FeedlyResponse
	testBlog := getTestBlog()
	testBlogs := []chase.Blogs{testBlog}
	testResponse.FeedTitle = "Test Title"
	testResponse.Items = testBlogs
	json, _ := json.Marshal(testResponse)
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(json)
	}))
}

func getTestBlog() chase.Blogs {
	var testBlog chase.Blogs
	testBlog.OriginID = "http://example.com"
	testBlog.Title = "Hello World"
	return testBlog
}
