package response

import (
	"errors"
	"fmt"
	"net/http"
)

/* HEADERS KEY */

const (
	contentType = "Content-Type"
	location    = "Location"
)

/* HEADERS VALUE */

const (
	applicationJson = "application/json"
)

func Json(w http.ResponseWriter, body []byte, code int) (int, error) {
	w.Header().Add(contentType, applicationJson)
	return response(w, body, code)
}

func JsonWithError(w http.ResponseWriter, err error, code int) (int, error) {
	if err == nil {
		return 0, errors.New("err cannot be nil")
	}
	bytes := formatMessageError(err.Error())
	return Json(w, bytes, code)
}

func JsonWithRedirect(w http.ResponseWriter, body []byte, redirect string, code int) (int, error) {
	w.Header().Add(location, redirect)
	return Json(w, body, code)
}

func formatMessageError(message string) []byte {
	return []byte(fmt.Sprintf("{\"message\": \"%s\"}", message))
}

func response(w http.ResponseWriter, body []byte, code int) (int, error) {
	w.WriteHeader(code)
	return w.Write(body)
}