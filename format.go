package rest

import (
	"fmt"
	"strings"
)

// defaultErrorMessage encapsulate an error in a json format
func defaultErrorMessage(err error) []byte {
	sanitize := strings.ReplaceAll(err.Error(), "\"", "")
	return []byte(fmt.Sprintf(`{"message":"%s"}`, sanitize))
}
