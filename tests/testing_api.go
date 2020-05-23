package tests

import (
	"chase/internal/chase"
	"reflect"
	"testing"
)

//This is the API used to facilitate tests

func constructBody() chase.RequestBody {
	return chase.RequestBody{"APT28"}
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
