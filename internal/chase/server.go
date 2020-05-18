package chase

import (
	"fmt"
	"net/http"
)

//Server is the handler function for the CHASE server
//that returns a slice of all APT related urls
func Server(writer http.ResponseWriter, request *http.Request) {
	apt := DecodeBody(request.Body)
	fmt.Fprint(writer, apt)
}
