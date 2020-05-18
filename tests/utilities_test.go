package tests

import (
	"chase/internal/chase"
	"io/ioutil"
	"strings"
	"testing"
)

//TestDecodeBody tests the DecodeBody utility function in internal/chase
func TestDecodeBody(t *testing.T) {
	t.Run("returns string from DecodeBody utility", func(t *testing.T) {
		reader := strings.NewReader("{{'APT':'APT28'}}")
		closer := ioutil.NopCloser(reader)
		got := chase.DecodeBody(closer)
		want := "APT28"
		assertEquals(t, got, want)
	})
}
