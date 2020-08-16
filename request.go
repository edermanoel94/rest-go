package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// CheckPathVariables see if any pathVariables match on params, if dont, add to a slice.
// Actually this works just on mux.Vars()
// TODO: make working in standard library
func CheckPathVariables(params map[string]string, pathVariables ...string) error {

	fields := make([]string, 0)

	for _, pathVariable := range pathVariables {
		if _, ok := params[pathVariable]; !ok {
			fields = append(fields, pathVariable)
		}
	}

	if len(fields) > 0 {
		return fmt.Errorf("params %s dont exists in the context", strings.Join(fields, ", "))
	}

	return nil
}

// GetPathVariable
func GetPathVariable(key string, params map[string]string) string {
	if param, ok := params[key]; !ok {
		return ""
	} else {
		return param
	}
}

// GetBody get the content of body on request and decode a pointer to a <T> to attach on body
func GetBody(reader io.ReadCloser, result interface{}) error {
	return decodeReader(reader, result)
}

func decodeReader(reader io.ReadCloser, result interface{}) error {

	defer reader.Close()

	if err := json.NewDecoder(reader).Decode(result); err != nil {
		return fmt.Errorf("couldn't decode: %v", err)
	}

	return nil
}
