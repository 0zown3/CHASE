package chase

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

//DecodeBody decodes a standard POST request body to retrieve the designated APT group
func DecodeBody(requestBody io.ReadCloser) string {
	var body RequestBody
	json.NewDecoder(requestBody).Decode(&body)
	return body.APT
}

//GetAliases returns an array of strings
func GetAliases(apt string) []string {
	aptMappings := GetAPTMappings()
	return aptMappings[apt].Groups
}

//GetAPTMappings returns an APTMappings interface
func GetAPTMappings() APTMappings {
	var aptMappings APTMappings
	jsonFile, _ := os.Open("../../config/apt_mappings.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &aptMappings)
	return aptMappings
}
