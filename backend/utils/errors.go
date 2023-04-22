package utils

import "net/http"

func HandleError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Write([]byte(message))
}
