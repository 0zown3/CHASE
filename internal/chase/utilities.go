package chase

import (
	"encoding/json"
	"io"
)

//DecodeBody decodes a standard POST request body to retrieve the designated APT group
func DecodeBody(requestBody io.ReadCloser) string {
	var body RequestBody
	json.NewDecoder(requestBody).Decode(&body)
	return body.APT
}
