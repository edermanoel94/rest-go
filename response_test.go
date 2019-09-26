package rest_test

import (
	"encoding/json"
	"errors"
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
type sliceError []int
type mapError map[string]interface{}

func (m mapError) Error() string {
	bytes, _ := json.Marshal(&m)
	return string(bytes)
}
func (c customError) Error() string {
	bytes, _ := json.Marshal(&c)
	return string(bytes)
}
func (s sliceError) Error() string {
	bytes, _ := json.Marshal(&s)
	return string(bytes)
}

// TESTS

func TestContent(t *testing.T) {

	testCases := []struct {
		description string
		payload     []byte
		statusCode  int
		isValidJson bool
	}{
		{"should serialize message in bytes and send statusCode",
			[]byte("{\"name\": \"cale\"}"), http.StatusOK, true},
		{"should not valid a json", []byte("{\"}"), http.StatusBadRequest, false},
		{"should send a nil in body of content", nil, http.StatusOK, false},
	}

	for _, tc := range testCases {

		t.Run(tc.description, func(t *testing.T) {

			recorder := httptest.NewRecorder()

			rest.Content(recorder, tc.payload, tc.statusCode)

			result := recorder.Result()

			defer result.Body.Close()

			payloadReceived, err := ioutil.ReadAll(result.Body)

			if err != nil {
				t.Fatalf("cannot read recorder: %v", err)
			}

			assert.Equal(t, tc.isValidJson, json.Valid(payloadReceived))
			assert.Equal(t, len(tc.payload), len(payloadReceived))
			assert.Equal(t, tc.statusCode, result.StatusCode)
			assert.Equal(t, "application/json", result.Header.Get("Content-Type"))
		})
	}
}

func TestError(t *testing.T) {

	testCases := []struct {
		description string
		err         error
		isValidJson bool
	}{
		{"should send a message of error with a status code", errors.New("not found"), true},
		{"should ignore extra quotes and valid json", errors.New("\"not found'"), true},
		{"should send a custom struct error message which implements error interface", customError{
			Description: "cannot found",
			Code:        "001",
		}, true},
		{"should send a custom map error message which implements error interface", mapError{"message": "error"}, true},
		{"should send a custom slice error message which implements error interface", sliceError{1, 2, 3}, true},
	}

	for _, tc := range testCases {

		t.Run(tc.description, func(t *testing.T) {

			recorder := httptest.NewRecorder()

			rest.Error(recorder, tc.err, http.StatusNotFound)

			result := recorder.Result()

			defer result.Body.Close()

			payloadReceived, err := ioutil.ReadAll(result.Body)

			if err != nil {
				t.Fatalf("cannot read recorder: %v", err)
			}

			assert.Equal(t, tc.isValidJson, json.Valid(payloadReceived))

			assert.Contains(t, string(payloadReceived), tc.err.Error())
		})
	}

	// TODO: add more tests with custom error on pointer
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

			payloadReceived, err := ioutil.ReadAll(result.Body)

			if err != nil {
				t.Fatal(err)
			}

			assert.True(t, json.Valid(payloadReceived))

			assert.Contains(t, string(payloadReceived), tc.contains)
		})
	}
}
