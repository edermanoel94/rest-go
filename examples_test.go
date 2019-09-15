package rest_test

import (
	"encoding/json"
	"errors"
	"github.com/edermanoel94/rest-go"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
)

func ExampleMarshalled() {
	product := struct {
		Name  string  `json:"name"`
		Price float32 `json:"price"`
	}{
		Name:  "Smart TV",
		Price: 150.20,
	}

	recorder := httptest.NewRecorder()

	rest.Marshalled(recorder, &product, http.StatusOK)

	result := recorder.Result()

	defer result.Body.Close()

	_, _ = io.Copy(os.Stdout, result.Body)

	// Output: {"name":"Smart TV","price":150.2}
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

	rest.Content(recorder, bytes, http.StatusOK)

	result := recorder.Result()

	defer result.Body.Close()

	_, _ = io.Copy(os.Stdout, result.Body)

	// Output: {"name":"Smart TV","price":100}
}

func ExampleError() {

	err := errors.New("cannot create product")

	recorder := httptest.NewRecorder()

	rest.Error(recorder, err, http.StatusOK)

	result := recorder.Result()

	defer result.Body.Close()

	_, _ = io.Copy(os.Stdout, result.Body)

	// Output: {"message":"cannot create product"}
}

func ExampleGetBody() {

}

func ExampleCheckPathVariables() {

}
