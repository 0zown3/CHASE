package chase

import (
	"encoding/json"
	"io"
)

//DecodeBody decodes the request body
func DecodeBody(requestBody io.ReadCloser) string {
	var body Request
	json.NewDecoder(requestBody).Decode(&body)
	return body.Token
}
