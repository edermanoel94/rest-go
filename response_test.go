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
}

type customErrorWithoutJson struct {
	Description string `json:"description"`
}

func (c customError) Error() string {
	bytes, _ := json.Marshal(&c)
	return string(bytes)
}

func (c customErrorWithoutJson) Error() string {
	return fmt.Sprintf("description: %s", c.Description)
}

// TESTS

func TestContent(t *testing.T) {

	testCases := []struct {
		description string
		payload     []byte
		statusCode  int
	}{
		{"should serialize message in bytes and send statusCode",
			[]byte("{\"name\": \"cale\"}"), http.StatusOK},
		{"should send a nil in body of content", nil, http.StatusOK},
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

			assert.Equal(t, len(tc.payload), len(payloadReceived))
			assert.Equal(t, tc.statusCode, result.StatusCode)
			assert.Equal(t, "application/json", result.Header.Get("Content-Type"))
		})
	}
}

func TestError(t *testing.T) {

	testCases := []struct {
		description    string
		err            error
		errStringGiven string
	}{
		{"should given a message of error", errors.New("not found"), "not found"},
		{"should ignore extra quotes", errors.New("\"not f\"ound"), "not found"},
		{"should send a custom struct error message which implements error interface",
			customError{Description: "not found"}, "{\"description\":\"not found\"}"},
		{"should send a custom struct which implements error interface but not use json.Marshal",
			customErrorWithoutJson{Description: "cannot found"}, "{\"message\":\"not a valid json\"}"},
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

			assert.Contains(t, string(payloadReceived), tc.errStringGiven)
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
		{"should marshal to a 0", 0, "0"},
		{"should marshal to a nil", nil, "null"},
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

			assert.Contains(t, string(payloadReceived), tc.contains)
		})
	}
}
