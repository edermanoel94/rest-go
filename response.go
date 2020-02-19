package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
)

var (
	ErrNotValidJson = errors.New("not a valid json")
)

// Response send slice of bytes to respond json
func Response(w http.ResponseWriter, body []byte, code int) (int, error) {
	if !json.Valid(body) {
		return response(w, defaultJsonErrorMessage(ErrNotValidJson), http.StatusInternalServerError)
	}
	return response(w, body, code)
}

// Marshalled use pointer to marshall and respond json
func Marshalled(w http.ResponseWriter, v interface{}, code int) (int, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return Error(w, err, http.StatusInternalServerError)
	}
	return Response(w, bytes, code)
}

// Error send a error to respond json, can send a non-struct which implements error.
func Error(w http.ResponseWriter, err error, code int) (int, error) {

	var errBytes []byte

	switch typeOf := reflect.TypeOf(err); typeOf.Kind() {
	case reflect.Ptr:
		errBytes = defaultJsonErrorMessage(err)
	default:
		errBytes = []byte(err.Error())
		return Response(w, errBytes, http.StatusInternalServerError)
	}

	return Response(w, errBytes, code)
}

func response(w http.ResponseWriter, body []byte, code int) (int, error) {
	w.Header().Set(contentType, applicationJson)
	w.WriteHeader(code)
	return w.Write(body)
}
