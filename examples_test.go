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
	"strings"
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

	rest.Response(recorder, bytes, http.StatusOK)

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

	product := struct {
		Name string `json:"name"`
	}{}

	readerCloser := ioutil.NopCloser(strings.NewReader("{\"name\": \"eder\"}"))

	err := rest.GetBody(readerCloser, &product)

	if err != nil {
		// do stuff with error
	}

	// Output:
}

func ExampleCheckPathVariables() {

	params := make(map[string]string)

	params["id"] = "21321423"

	err := rest.CheckPathVariables(params, "id")

	if err != nil {
		// do stuff with error
	}

	// Output:
}
