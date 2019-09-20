package rest_test

import (
	"fmt"
	"github.com/edermanoel94/rest-go"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func TestCheckPathVariables(t *testing.T) {

	t.Run("should send three path variables", func(t *testing.T) {

		params := make(map[string]string)

		params["name"] = "eder"
		params["last_name"] = "costa"

		err := rest.CheckPathVariables(params, "name", "last_name")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, params["name"], "eder")
		assert.Equal(t, params["last_name"], "costa")
	})

	t.Run("should failed if not match with params", func(t *testing.T) {

		params := make(map[string]string)

		params["name"] = "eder"
		params["last_name"] = "Ede$1"

		err := rest.CheckPathVariables(params, "namE", "lastname")

		if err == nil {
			t.Fatalf("expect params %s dont exists in the context", params)
		}

		assert.Contains(t, err.Error(), "namE")
		assert.Contains(t, err.Error(), "lastname")
	})
}

func TestGetBody(t *testing.T) {

	t.Run("should get a body on reader and marshal", func(t *testing.T) {

		result := struct {
			Name string `json:"name"`
		}{}

		readerCloser := ioutil.NopCloser(strings.NewReader("{\"name\": \"eder\"}"))

		err := rest.GetBody(readerCloser, &result)

		if err != nil {
			t.Fatalf("%v", err)
		}

		assert.Equal(t, "eder", result.Name, fmt.Sprintf("expected: %s, got: %s", "eder", result.Name))
	})

	t.Run("should not unmarshal if send a nil on result", func(t *testing.T) {

		reader := strings.NewReader("{\"name\": \"eder\"}")

		readerCloser := ioutil.NopCloser(reader)

		err := rest.GetBody(readerCloser, nil)

		if err == nil {
			t.Fatal(err)
		}

		assert.Contains(t, err.Error(), "couldn't unmarshal")
	})

	t.Run("should not unmarshal if send a Reader with empty string", func(t *testing.T) {

		result := struct {
			Name string `json:"name"`
		}{}

		reader := strings.NewReader("")

		readerCloser := ioutil.NopCloser(reader)

		err := rest.GetBody(readerCloser, result)

		if err == nil {
			t.Fatal(err)
		}

		assert.Contains(t, err.Error(), "couldn't unmarshal")
	})
}
