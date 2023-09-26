package exceptions

import "net/http"

func JSONResponseException(err error, responseWriter http.ResponseWriter) bool {
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return true
	}
	return false
}
