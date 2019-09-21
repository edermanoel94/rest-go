package main

import (
	"encoding/json"
	"errors"
	"github.com/edermanoel94/rest-go"
	"github.com/gorilla/mux"
	"github.com/nanobox-io/golang-scribble"
	"github.com/rs/cors"
	"github.com/teris-io/shortid"
	"log"
	"net/http"
)

type product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
}

type handler struct {
	db *scribble.Driver
}

const (
	collection = "product"
)

func main() {

	db, _ := scribble.New("./db", &scribble.Options{})

	handler := handler{db: db}

	router := mux.NewRouter()

	router.HandleFunc("/", handler.saveHandler).Methods(http.MethodPost)
	router.HandleFunc("/", handler.listHandler).Methods(http.MethodGet)
	router.HandleFunc("/{resource}", handler.detailHandler).Methods(http.MethodGet)
	router.HandleFunc("/{resource}", handler.deleteHandler).Methods(http.MethodDelete)
	router.HandleFunc("/{resource}", handler.updateHandler).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe("0.0.0.0:80", cors.AllowAll().Handler(router)))
}

//
func (h *handler) saveHandler(w http.ResponseWriter, r *http.Request) {

	product := product{}

	err := rest.GetBody(r.Body, &product)

	if err != nil {
		rest.Error(w, err, http.StatusBadRequest)
		return
	}

	id, _ := shortid.Generate()

	product.ID = id
	err = h.db.Write(collection, product.ID, &product)

	if err != nil {
		rest.Error(w, err, http.StatusInternalServerError)
		return
	}

	rest.Marshalled(w, &product, http.StatusCreated)
}

//
func (h *handler) listHandler(w http.ResponseWriter, r *http.Request) {

	records, _ := h.db.ReadAll(collection)
	products := make([]product, 0)

	for _, record := range records {
		product := product{}
		json.Unmarshal([]byte(record), &product)
		products = append(products, product)
	}

	rest.Marshalled(w, &products, http.StatusOK)
}

//
func (h *handler) detailHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	err := rest.CheckPathVariables(params, "resource")

	if err != nil {
		rest.Error(w, err, http.StatusBadRequest)
		return
	}

	product := product{}

	err = h.db.Read(collection, params["resource"], &product)

	if err != nil {
		rest.Error(w, err, http.StatusInternalServerError)
		return
	}

	rest.Marshalled(w, &product, http.StatusOK)
}

//
func (h *handler) deleteHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	err := rest.CheckPathVariables(params, "resource")

	if err != nil {
		rest.Error(w, err, http.StatusBadRequest)
		return
	}

	err = h.db.Delete(collection, params["resource"])

	if err != nil {
		rest.Error(w, err, http.StatusInternalServerError)
		return
	}

	rest.Content(w, nil, http.StatusOK)
}

//
func (h *handler) updateHandler(w http.ResponseWriter, r *http.Request) {
	rest.Error(w, errors.New("not implemented"), http.StatusNotImplemented)
}
