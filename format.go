package rest

import "fmt"

// defaultErrorMessage encapsulate an error in a json format
func defaultErrorMessage(err error) []byte {
	return []byte(fmt.Sprintf("{\"message\":\"%s\"}", err.Error()))
}
