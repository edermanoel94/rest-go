package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

func Content(w http.ResponseWriter, body []byte, code int) (int, error) {
	w.Header().Add(contentType, applicationJson)
	return response(w, body, code)
}

// TODO: check if body match with struct
func Marshalled(w http.ResponseWriter, v interface{}, code int) (int, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return Error(w, err, http.StatusBadRequest)
	}
	return Content(w, bytes, code)
}

func Error(w http.ResponseWriter, err error, code int) (int, error) {

	var bytes []byte

	switch typeOf := reflect.TypeOf(err); typeOf.Kind() {
	case reflect.Struct:
		bytes, err = json.Marshal(err)
		if err != nil {
			return Content(w, formatMessageError(err.Error()), http.StatusInternalServerError)
		}
	default:
		bytes = formatMessageError(err.Error())
	}

	return Content(w, bytes, code)
}

func Location(w http.ResponseWriter, body []byte, url string, code int) (int, error) {
	w.Header().Add(location, url)
	return Content(w, body, code)
}

func formatMessageError(message string) []byte {
	return []byte(fmt.Sprintf("{\"message\": \"%s\"}", message))
}

func response(w http.ResponseWriter, body []byte, code int) (int, error) {
	w.WriteHeader(code)
	return w.Write(body)
}
