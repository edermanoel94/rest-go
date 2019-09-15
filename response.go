package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
)

var (
	ErrIsNil = errors.New("error cannot be nil")
)

func Content(w http.ResponseWriter, body []byte, code int) (int, error) {
	w.Header().Add(contentType, applicationJson)
	return response(w, body, code)
}

func Marshalled(w http.ResponseWriter, v interface{}, code int) (int, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return Error(w, err, http.StatusBadRequest)
	}
	return Content(w, bytes, code)
}

func Error(w http.ResponseWriter, err error, code int) (int, error) {

	if err == nil {
		return Content(w, defaultErrorMessage(ErrIsNil), http.StatusInternalServerError)
	}

	var bytes []byte

	switch typeOf := reflect.TypeOf(err); typeOf.Kind() {
	case reflect.Struct:
		bytes, err = json.Marshal(err)
		if err != nil {
			return Content(w, defaultErrorMessage(err), http.StatusInternalServerError)
		}
	default:
		bytes = defaultErrorMessage(err)
	}

	return Content(w, bytes, code)
}

func response(w http.ResponseWriter, body []byte, code int) (int, error) {
	w.WriteHeader(code)
	return w.Write(body)
}
