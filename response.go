package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
)

var (
	// ErrIsNil is when send a nil error
	ErrIsNil = errors.New("error cannot be nil")
)

// Content send slice of bytes to respond json
func Content(w http.ResponseWriter, body []byte, code int) (int, error) {
	w.Header().Add(contentType, applicationJson)
	return response(w, body, code)
}

// Marshalled use pointer to marshall and respond json
func Marshalled(w http.ResponseWriter, v interface{}, code int) (int, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return Error(w, err, http.StatusBadRequest)
	}
	return Content(w, bytes, code)
}

// Error send a error to respond json, can send a non-struct which implements error
// and stringify.
func Error(w http.ResponseWriter, err error, code int) (int, error) {

	var bytes []byte

	switch typeOf := reflect.TypeOf(err); typeOf.Kind() {
	case reflect.Ptr:
		bytes = defaultErrorMessage(err)
	default:
		return Content(w, []byte(err.Error()), http.StatusInternalServerError)
	}

	return Content(w, bytes, code)
}

func response(w http.ResponseWriter, body []byte, code int) (int, error) {
	w.WriteHeader(code)
	return w.Write(body)
}
