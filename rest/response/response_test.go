package response_test

import (
	"errors"
	"github.com/edermanoel94/cale/rest/response"
	"github.com/gobuffalo/httptest"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestResponse_Json(t *testing.T) {

	recorder := httptest.NewRecorder()

	actualBody := []byte("oi")

	actualStatusCode := http.StatusOK

	_, _ = response.Json(recorder, actualBody, actualStatusCode)

	result := recorder.Result()

	bytes, err := ioutil.ReadAll(result.Body)

	if err != nil {
		t.Fatalf("cannot read recorder: %v", err)
	}

	defer result.Body.Close()

	if len(actualBody) != len(bytes) {
		t.Fatalf("size of slice of bytes is different")
	}

	if actualStatusCode != result.StatusCode {
		t.Fatalf("got status %d, but given %d", actualStatusCode, result.StatusCode)
	}

	contentType := "Content-Type"

	if result.Header.Get(contentType) != "application/json" {
		t.Fatalf("should be application/json, got: %s", result.Header.Get(contentType))
	}
}

func TestJsonWithError(t *testing.T) {

	t.Run("should respond with json", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		actualError := errors.New("error")

		actualStatusCode := http.StatusNotFound

		_, _ = response.JsonWithError(recorder, actualError, actualStatusCode)

		result := recorder.Result()

		bytes, err := ioutil.ReadAll(result.Body)

		if err != nil {
			t.Fatalf("cannot read recorder: %v", err)
		}

		defer result.Body.Close()

		if actualError.Error() == string(bytes) {
			t.Fatalf("")
		}
	})

	t.Run("should not respond", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		actualError := errors.New("error")

		actualStatusCode := http.StatusNotFound

		_, _ = response.JsonWithError(recorder, actualError, actualStatusCode)

		result := recorder.Result()

		bytes, err := ioutil.ReadAll(result.Body)

		if err != nil {
			t.Fatalf("cannot read recorder: %v", err)
		}

		defer result.Body.Close()

		if actualError.Error() == string(bytes) {
			t.Fatalf("")
		}
	})
}

func TestJsonWithRedirect(t *testing.T) {

	recorder := httptest.NewRecorder()

	actualBody := []byte("oi")
	actualStatusCode := http.StatusOK
	actualRedirect := "http://localhost:8080"

	_, _ = response.JsonWithRedirect(recorder, actualBody, actualRedirect, actualStatusCode)

	result := recorder.Result()

	_, err := ioutil.ReadAll(result.Body)

	if err != nil {
		t.Fatalf("cannot read recorder: %v", err)
	}

	defer result.Body.Close()

	location := "Location"

	if result.Header.Get(location) != actualRedirect {
		t.Fatalf("should be application/json, got: %s", result.Header.Get(location))
	}
}
