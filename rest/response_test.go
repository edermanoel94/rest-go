package rest_test

import (
	"encoding/json"
	"errors"
	"github.com/edermanoel94/cale/rest"
	"github.com/gobuffalo/httptest"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestResponse_Json(t *testing.T) {

	recorder := httptest.NewRecorder()

	actualBody := []byte("{\"name\": \"cale\"}")

	actualStatusCode := http.StatusOK

	_, _ = rest.Json(recorder, actualBody, actualStatusCode)

	result := recorder.Result()

	defer result.Body.Close()

	bytes, err := ioutil.ReadAll(result.Body)

	if err != nil {
		t.Fatalf("cannot read recorder: %v", err)
	}

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

	t.Run("should send a signal error", func(t *testing.T) {

		actualError := errors.New("error")
		actualStatusCode := http.StatusNotFound

		recorder := httptest.NewRecorder()

		_, _ = rest.JsonWithError(recorder, actualError, actualStatusCode)

		result := recorder.Result()

		defer result.Body.Close()

		bytes, err := ioutil.ReadAll(result.Body)

		if err != nil {
			t.Fatalf("cannot read recorder: %v", err)
		}

		content := map[string]string{}

		err = json.Unmarshal(bytes, &content)

		if err != nil {
			t.Fatalf("couldn't unmarshal: %v", err)
		}

		if actualError.Error() != content["message"] {
			t.Fatalf("expected: %s, got: %s", actualError.Error(), content["message"])
		}
	})
}

func TestJsonWithRedirect(t *testing.T) {

	recorder := httptest.NewRecorder()

	actualBody := []byte("{\"name\": \"cale\"}")
	actualStatusCode := http.StatusOK
	actualRedirect := "http://localhost:8080/tenant/LASA"

	_, _ = rest.JsonWithRedirect(recorder, actualBody, actualRedirect, actualStatusCode)

	result := recorder.Result()

	defer result.Body.Close()

	_, err := ioutil.ReadAll(result.Body)

	if err != nil {
		t.Fatalf("cannot read recorder: %v", err)
	}

	location := "Location"

	headerLocation := result.Header.Get(location)

	if actualRedirect != headerLocation {
		t.Fatalf("expected a redirect to %s, got: %s", headerLocation, headerLocation)
	}
}
