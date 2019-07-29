package rest

import (
	"fmt"
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
		return fmt.Errorf("params %s dont exists", strings.Join(fields, ", "))
	}

	return nil
}
