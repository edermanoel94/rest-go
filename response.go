package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Json(w http.ResponseWriter, body []byte, code int) (int, error) {
	w.Header().Add(contentType, applicationJson)
	return response(w, body, code)
}

func Marshalled(w http.ResponseWriter, v interface{}, code int) (int, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return Error(w, err, http.StatusBadRequest)
	}
	return Json(w, bytes, code)
}

func Error(w http.ResponseWriter, err error, code int) (int, error) {
	if err == nil {
		return 0, fmt.Errorf("err cannot be null")
	}
	bytes := formatMessageError(err.Error())
	return Json(w, bytes, code)
}

func Redirect(w http.ResponseWriter, body []byte, redirect string, code int) (int, error) {
	w.Header().Add(location, redirect)
	return Json(w, body, code)
}

// TODO: Customize
func formatMessageError(message string) []byte {
	return []byte(fmt.Sprintf("{\"message\": \"%s\"}", message))
}

//func formatCustomMessageError(err error) []byte {
//	value := reflect.ValueOf(err)
//	switch value.Kind() {
//	case reflect.Struct:
//		formatMessageError(err.Error())
//	default:
//
//	}
//}

func response(w http.ResponseWriter, body []byte, code int) (int, error) {
	w.WriteHeader(code)
	return w.Write(body)
}

type CustomErrorJson struct {
	Message string
	Path    string
	Code    int
}

func (e CustomErrorJson) Error() string {
	bytes, _ := json.Marshal(&e)
	return string(bytes)
}
