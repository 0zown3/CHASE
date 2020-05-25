package tests

import (
	"bytes"
	"chase/internal/chase"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"testing"
)

//This is the API used to facilitate tests

func constructBody() chase.Request {
	return chase.Request{"test123"}
}

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

func getRequestBody() io.ReadCloser {
	requestBody := constructBody()
	jsonBody, _ := json.Marshal(requestBody)
	request, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))
	return request.Body
}
