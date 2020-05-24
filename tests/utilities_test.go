package tests

import (
	"chase/internal/chase"
	"reflect"
	"testing"
)

//TestDecodeBody tests the DecodeBody utility function in internal/chase
func TestDecodeBody(t *testing.T) {
	t.Run("returns string from DecodeBody utility", func(t *testing.T) {
		requestBody := getRequestBody()
		got := chase.DecodeBody(requestBody)
		want := "APT28"
		assertEquals(t, got, want)
	})
}

//TestGetAliases tests the return type of chase.GetAliases (we want []string)
func TestGetAliases(t *testing.T) {
	t.Run("returns array of aliases from apt_mappings.json", func(t *testing.T) {
		var groups []string
		var testGroups []string
		var apt string
		apt = "APT28"
		groups = chase.GetAliases(apt)
		got := reflect.TypeOf(groups)
		want := reflect.TypeOf(testGroups)
		assertType(t, got, want)
	})
}

//TestGetAPTMappings tests the chase.GetAPTMappings utility
func TestGetAPTMappings(t *testing.T) {
	t.Run("tests chase.ReadFile utility", func(t *testing.T) {
		var testMappings chase.APTMappings
		got := chase.GetAPTMappings()
		want := reflect.TypeOf(testMappings)
		assertType(t, reflect.TypeOf(got), want)
	})
}
