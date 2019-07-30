package rest

import (
	"fmt"
	"net/http"
)

// Headers keys
const (
	contentType = "Content-Type"
	location    = "Location"
)

// Headers values
const (
	applicationJson = "application/json"
)

func Json(w http.ResponseWriter, body []byte, code int) (int, error) {
	w.Header().Add(contentType, applicationJson)
	return response(w, body, code)
}

func JsonWithError(w http.ResponseWriter, err error, code int) (int, error) {
	if err == nil {
		return 0, fmt.Errorf("error cannot be null")
	}
	bytes := formatMessageError(err.Error())
	return Json(w, bytes, code)
}

func JsonWithRedirect(w http.ResponseWriter, body []byte, redirect string, code int) (int, error) {
	w.Header().Add(location, redirect)
	return Json(w, body, code)
}

// TODO: Customize
func formatMessageError(message string) []byte {
	return []byte(fmt.Sprintf("{\"message\": \"%s\"}", message))
}

func response(w http.ResponseWriter, body []byte, code int) (int, error) {
	w.WriteHeader(code)
	return w.Write(body)
}
