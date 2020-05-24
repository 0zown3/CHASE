package chase

import (
	"net/http"
)

//Server is the handler function for the CHASE server
func Server(writer http.ResponseWriter, request *http.Request) {
	Chase(request.Body)
}
