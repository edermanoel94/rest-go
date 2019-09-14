Rest GO - Helpful library for Rest API
================================

A package that provide many helpful methods for working with rest api.

And if u are tired of written, this:

```go

func SomeHandler(w http.ResponseWriter, r *http.Request) {

    w.Header().Add("Content-Type", "application/json")

    product := &product{"Smart TV", 50.00}
    
    bytes, err := json.Marshal(product)
    
    if err != nil {
    	w.WriteHeader(http.StatusInternalServerError)
        // super bad!
        message := fmt.Sprintf("{\"message\": \"%s\"}", err.Error())
    	w.Write([]byte(message))
	    return
    }
    
    w.WriteHeader(http.StatusOk)
    w.Write(bytes)
}
```

Get started:

  * Install rest GO with [one line of code](#installation)


[`rest`](http://godoc.org/github.com/edermanoel94/rest-go "API documentation") package
-------------------------------------------------------------------------------------------

The `rest` package provides some helpful methods that allow you to write better rest api in GO.

  * Allows for very readable code

See it in action:

```go
package yours

import (
    "encoding/json"
    "github.com/edermanoel94/rest-go"
    "net/http"
)

type product struct {
    Name  string `json:"name"`
    Price float32 `json:"price"`
}

func SomeHandler(w http.ResponseWriter, r *http.Request) {
	
    product := &product{"Smart TV", 50.00}

    // rest.Marshalled marshall the struct and respond json.
    rest.Marshalled(w, product, http.StatusOK)
}
```

Working with body request on methods: [PUT, PATCH, POST]:

```go
package yours

import (
    "encoding/json"
    "github.com/edermanoel94/rest-go"
    "net/http"
)

type product struct {
    Name  string `json:"name"`
    Price float32 `json:"price"`
}

// [POST] body: {"name": "eder", "price": 20.00}
func SomePostHandler(w http.ResponseWriter, r *http.Request) {

    product := product{}

    rest.GetBody(r.Body, &product)

    // Save/Update Whatever ur want to do with body
}
```

Working with Path variables:

```go
package yours

import (
    "encoding/json"
    "github.com/edermanoel94/rest-go"
    "net/http"
)

type product struct {
    Name  string `json:"name"`
    Price float32 `json:"price"`
}

// [POST] body: {"name": "eder", "price": 20.00}
func SomePostHandler(w http.ResponseWriter, r *http.Request) {

    product := product{}

    rest.GetBody(r.Body, &product)

    // Save/Update Whatever ur want to do with body
}    
```

Working with [`mux`](https://github.com/gorilla/mux "API documentation")

```go
package yours

import (
    "encoding/json"
    "github.com/edermanoel94/rest-go"
    "net/http"
)

type product struct {
    Name  string `json:"name"`
    Price float32 `json:"price"`
}

// [GET] url: /product/{id} body: {"name": "eder", "price": 20.00}
func SomePostHandler(w http.ResponseWriter, r *http.Request) {
	
    params := mux.Vars(r)
    
    rest.CheckPathVariables(params, "id")
}
```

TODO List
=========

- [ ] Working with pagination and test
- [ ] Working with custom errors
- [x] Test unit for request
- [x] Test unit for response

Installation
============

To install, use `go get`:

    go get github.com/edermanoel94/rest-go


Contributing
============

Please feel free to submit issues, fork the repository and send pull requests!

------

License
=======

This project is licensed under the terms of the MIT license.
