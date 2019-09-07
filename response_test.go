package rest_test

import (
	"encoding/json"
	"errors"
	"github.com/edermanoel94/rest-go"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type CustomError struct {
	Description string `json:"description"`
	Code        string `json:"code"`
}

func (c CustomError) Error() string {
	bytes, _ := json.Marshal(&c)
	return string(bytes)
}

func TestContent(t *testing.T) {

	// TODO: make many jsons invalid to check

	t.Run("should serialize message in bytes and send statusCode", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		body := []byte("{\"name\": \"cale\"}")

		statusCode := http.StatusOK

		_, _ = rest.Content(recorder, body, statusCode)

		result := recorder.Result()

		defer result.Body.Close()

		bytes, err := ioutil.ReadAll(result.Body)

		if err != nil {
			t.Fatalf("cannot read recorder: %v", err)
		}

		if len(body) != len(bytes) {
			t.Fatalf("size of slice of bytes is different")
		}

		if statusCode != result.StatusCode {
			t.Fatalf("got status %d, but given %d", statusCode, result.StatusCode)
		}

		contentType := "Content-Type"

		if result.Header.Get(contentType) != "application/json" {
			t.Fatalf("should be application/json, got: %s", result.Header.Get(contentType))
		}
	})

	t.Run("should send a nil in message and send statusCode", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		var body []byte

		statusCode := http.StatusOK

		_, _ = rest.Content(recorder, body, statusCode)

		result := recorder.Result()

		defer result.Body.Close()

		bytes, err := ioutil.ReadAll(result.Body)

		if err != nil {
			t.Fatalf("cannot read recorder: %v", err)
		}

		if len(body) != len(bytes) {
			t.Fatalf("size of slice of bytes is different")
		}

		if statusCode != result.StatusCode {
			t.Fatalf("got status %d, but given %d", statusCode, result.StatusCode)
		}

		contentType := "Content-Type"

		if result.Header.Get(contentType) != "application/json" {
			t.Fatalf("should be application/json, got: %s", result.Header.Get(contentType))
		}
	})
}

func TestError(t *testing.T) {

	t.Run("should send a message of error with a status code 404", func(t *testing.T) {

		errorThrowed := errors.New("not found")
		statusCode := http.StatusNotFound

		recorder := httptest.NewRecorder()

		_, _ = rest.Error(recorder, errorThrowed, statusCode)

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

		if errorThrowed.Error() != content["message"] {
			t.Fatalf("expected: %s, got: %s", errorThrowed.Error(), content["message"])
		}

		if statusCode != result.StatusCode {
			t.Fatalf("expected: %d, got: %d", statusCode, result.StatusCode)
		}
	})

	t.Run("should send a custom error message", func(t *testing.T) {

		customError := CustomError{
			Description: "cannot found",
			Code:        "001",
		}
		statusCode := http.StatusNotFound

		recorder := httptest.NewRecorder()

		_, _ = rest.Error(recorder, customError, statusCode)

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

		if customError.Description != content["description"] {
			t.Fatalf("expected: %s, got: %s", customError.Description, content["description"])
		}

		if customError.Code != content["code"] {
			t.Fatalf("expected: %s, got: %s", customError.Code, content["code"])
		}

		if statusCode != result.StatusCode {
			t.Fatalf("expected: %d, got: %d", statusCode, result.StatusCode)
		}
	})
}

func TestLocation(t *testing.T) {

	t.Run("", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		body := []byte("{\"name\": \"cale\"}")
		statusCode := http.StatusOK
		redirect := "http://localhost:8080/"

		_, _ = rest.Location(recorder, body, redirect, statusCode)

		result := recorder.Result()

		defer result.Body.Close()

		_, err := ioutil.ReadAll(result.Body)

		if err != nil {
			t.Fatalf("cannot read recorder: %v", err)
		}

		location := "Location"

		headerLocation := result.Header.Get(location)

		if redirect != headerLocation {
			t.Fatalf("expected a redirect to %s, got: %s", headerLocation, headerLocation)
		}
	})
}

func ExampleContent() {

	product := struct {
		Name  string  `json:"name"`
		Price float32 `json:"price"`
	}{
		Name:  "Smart TV",
		Price: 100.00,
	}

	bytes, _ := json.Marshal(&product)

	recorder := httptest.NewRecorder()

	_, _ = rest.Content(recorder, bytes, http.StatusOK)

	result := recorder.Result()

	defer result.Body.Close()

	_, _ = io.Copy(os.Stdout, result.Body)

	// Output: {"name":"Smart TV","price":100}

}
