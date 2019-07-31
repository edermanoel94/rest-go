package rest

import (
	"encoding/json"
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

func JsonMarshalled(w http.ResponseWriter, v interface{}, code int) (int, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return JsonWithError(w, err, http.StatusBadRequest)
	}
	return Json(w, bytes, code)
}

func JsonWithError(w http.ResponseWriter, err error, code int) (int, error) {
	if err == nil {
		return 0, fmt.Errorf("err cannot be null")
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
