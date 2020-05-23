package tests

import (
	"bytes"
	"chase/internal/chase"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

//TestServer tests the response code returned by ChaseServer
func TestServer(t *testing.T) {
	t.Run("returns response from fetching related APT urls", func(t *testing.T) {
		requestBody := constructBody()
		jsonBody, _ := json.Marshal(requestBody)
		request, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))
		response := httptest.NewRecorder()
		chase.Server(response, request)
		assertStatus(t, response.Code, 200)
	})
}

//TODO: Need utility function to assert on response type (aka slice)
