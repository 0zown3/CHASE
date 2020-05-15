package tests

import (
	"chase/internal/chase"
	"net/http"
	"net/http/httptest"
	"testing"
)

//TestServer tests the response code returned by ChaseServer
func TestServer(t *testing.T) {
	t.Run("returns response from fetching related APT urls", func(t *testing.T) {
		//This nil needs to be replaced with appropriate body data.
		request, _ := http.NewRequest(http.MethodPost, "/", nil)
		response := httptest.NewRecorder()

		chase.Server(response, request)

		got := response.Code
		want := 200

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
