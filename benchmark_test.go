package rest_test

import (
	"github.com/edermanoel94/rest-go"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func BenchmarkGetBody(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result := struct {
			Name string `json:"name"`
		}{}
		err := rest.GetBody(generateReadCloser(), &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalled(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recorder := httptest.NewRecorder()
		result := struct {
			Name string `json:"name"`
		}{"Eder"}
		rest.Marshalled(recorder, &result, http.StatusInternalServerError)
	}
}

func generateReadCloser() io.ReadCloser {
	return ioutil.NopCloser(strings.NewReader("{\"name\": \"eder\"}"))
}
