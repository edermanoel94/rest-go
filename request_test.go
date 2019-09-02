package rest_test

import (
	"github.com/edermanoel94/rest-go"
	"io/ioutil"
	"strings"
	"testing"
)

func TestCheckPathVariables(t *testing.T) {

	t.Run("should", func(t *testing.T) {

		params := make(map[string]string)

		params["eder"] = "ok"
		params["eder"] = "qweqq"
		params["manoel"] = "qweqq"

		err := rest.CheckPathVariables(params, "eder", "manoel", "eder")

		if err != nil {
			t.Fatalf("%v: ", err)
		}
	})
}

func TestGetBody(t *testing.T) {

	t.Run("should get a body on reader and marshal", func(t *testing.T) {

		result := struct {
			Name string `json:"name"`
		}{}

		reader := strings.NewReader("{\"name\": \"eder\"}")

		readerCloser := ioutil.NopCloser(reader)

		err := rest.GetBody(readerCloser, &result)

		if err != nil {
			t.Fatalf("%v", err)
		}

		if result.Name != "eder" {
			t.Fatalf("expected: %s, got: %s", "eder", result.Name)
		}

	})
}
