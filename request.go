package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// CheckPathVariables see if any pathVariables match on params, if dont, add to a slice.
// Actually this works just on mux.Vars()
// Can use safelly!
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

// GetBody get the content of body on request and unmarshal a pointer to a <T> to attach on body
func GetBody(reader io.ReadCloser, result interface{}) error {

	bytes, err := ioutil.ReadAll(reader)

	defer reader.Close()

	if err != nil {
		return fmt.Errorf("couldn't read body of request: %v", err)
	}

	err = json.Unmarshal(bytes, result)

	if err != nil {
		return fmt.Errorf("couldn't unmarshal: %v", err)
	}

	return nil
}
