package rest_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/edermanoel94/rest-go"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type customError struct {
	Description string `json:"description"`
	Code        string `json:"code"`
}

func (c customError) Error() string {
	bytes, _ := json.Marshal(&c)
	return string(bytes)
}

func TestContent(t *testing.T) {

	// TODO: make many jsons invalid to check, WIP

	t.Run("should serialize message in bytes and send statusCode", func(t *testing.T) {

		payloadSend := []byte("{\"name\": \"cale\"}")
		statusCode := http.StatusOK

		recorder := httptest.NewRecorder()

		rest.Content(recorder, payloadSend, statusCode)

		result := recorder.Result()

		defer result.Body.Close()

		bytesReceived, err := ioutil.ReadAll(result.Body)

		if err != nil {
			t.Fatalf("cannot read recorder: %v", err)
		}

		assert.Equal(t, len(payloadSend), len(bytesReceived))
		assert.Equal(t, statusCode, result.StatusCode)
		assert.Equal(t, "application/json", result.Header.Get("Content-Type"))
	})

	t.Run("should send a nil in message and send statusCode", func(t *testing.T) {

		recorder := httptest.NewRecorder()

		var payloadSend []byte
		statusCode := http.StatusOK

		rest.Content(recorder, nil, statusCode)

		result := recorder.Result()

		defer result.Body.Close()

		bytesReceived, err := ioutil.ReadAll(result.Body)

		if err != nil {
			t.Fatalf("cannot read recorder: %v", err)
		}

		assert.Equal(t, len(payloadSend), len(bytesReceived))
		assert.Equal(t, statusCode, result.StatusCode)
		assert.Equal(t, "application/json", result.Header.Get("Content-Type"))
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

		payloadReceived, err := ioutil.ReadAll(result.Body)

		if err != nil {
			t.Fatalf("cannot read recorder: %v", err)
		}

		assert.Contains(t, string(payloadReceived), "not found")
		assert.Equal(t, statusCode, result.StatusCode)
	})

	t.Run("should send a custom error message", func(t *testing.T) {

		customError := customError{
			Description: "cannot found",
			Code:        "001",
		}
		statusCode := http.StatusNotFound

		recorder := httptest.NewRecorder()

		rest.Error(recorder, customError, statusCode)

		result := recorder.Result()

		defer result.Body.Close()

		payloadReceived, err := ioutil.ReadAll(result.Body)

		if err != nil {
			t.Fatalf("cannot read recorder: %v", err)
		}

		fmt.Println(string(payloadReceived))
	})
}

// TODO: error if not send a location
func TestLocation(t *testing.T) {

	t.Run("should send a message with header `Location` and url", func(t *testing.T) {

		recorder := httptest.NewRecorder()

		payloadSend := []byte("{\"name\": \"cale\"}")
		statusCode := http.StatusOK
		redirect := "http://localhost:8080/"

		_, _ = rest.Location(recorder, payloadSend, redirect, statusCode)

		result := recorder.Result()

		defer result.Body.Close()

		_, err := ioutil.ReadAll(result.Body)

		if err != nil {
			t.Fatalf("cannot read recorder: %v", err)
		}

		assert.Equal(t, redirect, result.Header.Get("Location"))
		assert.Equal(t, statusCode, result.StatusCode)
	})
}

func TestMarshalled(t *testing.T) {

	t.Run("should marshal struct correctly", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		exampleToMarshal := struct {
			Name string `json:"name"`
		}{"Eder"}

		rest.Marshalled(recorder, &exampleToMarshal, http.StatusInternalServerError)

		result := recorder.Result()

		defer result.Body.Close()

		body, err := ioutil.ReadAll(result.Body)

		if err != nil {
			t.Fatal(err)
		}

		assert.Contains(t, string(body), "name")
		assert.Contains(t, string(body), "Eder")
	})

	t.Run("should not marshal if is a non-struct", func(t *testing.T) {

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
