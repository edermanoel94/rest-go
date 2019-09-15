package rest_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/edermanoel94/rest-go"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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

	t.Run("should send a message of error with a status code", func(t *testing.T) {

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

	t.Run("should send a nil error and given a `ErrIsNil`", func(t *testing.T) {

		var errorThrowed error
		statusCode := http.StatusNotFound

		recorder := httptest.NewRecorder()

		rest.Error(recorder, errorThrowed, statusCode)

		result := recorder.Result()

		defer result.Body.Close()

		payloadReceived, err := ioutil.ReadAll(result.Body)

		if err != nil {
			t.Fatalf("cannot read recorder: %v", err)
		}

		assert.Contains(t, string(payloadReceived), rest.ErrIsNil.Error())
		assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
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

func TestMarshalled(t *testing.T) {

	testCases := []struct {
		description string
		actual      interface{}
		contains    string
	}{
		{"should marshal struct correctly", struct {
			Name string `json:"name"`
		}{"Eder"}, "Eder"},
		{"should marshal to a 0 if is a pointer to int", 0, "0"},
		{"should not marshal to a null if is a nil", nil, "null"},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			recorder := httptest.NewRecorder()

			rest.Marshalled(recorder, &tc.actual, http.StatusInternalServerError)

			result := recorder.Result()

			defer result.Body.Close()

			body, err := ioutil.ReadAll(result.Body)

			if err != nil {
				t.Fatal(err)
			}

			assert.Contains(t, string(body), tc.contains)
		})

	}
}
