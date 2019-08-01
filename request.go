package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// CheckPathVariables see if any pathVariables match on params, if dont,
// just add to a slice and return a error
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

// GetBodyRequest get the content body of request and unmarshal a reference to a struct
func GetBodyOnRequest(r *http.Request, result interface{}) error {

	bytes, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	if err != nil {
		return fmt.Errorf("couldn't read body of request: %v", err)
	}

	err = json.Unmarshal(bytes, result)

	if err != nil {
		return fmt.Errorf("couldn't unmarshal: %v", err)
	}

	return nil
}
