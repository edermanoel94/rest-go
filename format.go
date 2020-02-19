package rest

import (
	"fmt"
	"strings"
)

// defaultJsonErrorMessage encapsulate an error in a json format
func defaultJsonErrorMessage(err error) []byte {
	sanitize := strings.ReplaceAll(err.Error(), "\"", "")
	return []byte(fmt.Sprintf(`{"message":"%s"}`, sanitize))
}
