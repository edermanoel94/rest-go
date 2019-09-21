Rest GO - Helpful library for Rest API
======================================

[![Build Status](https://travis-ci.org/edermanoel94/rest-go.svg?branch=master)](https://travis-ci.org/edermanoel94/rest-go)
[![eder](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/edermanoel94/rest-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/edermanoel94/rest-go)](https://goreportcard.com/report/github.com/edermanoel94/rest-go)

A package that provide many helpful methods for working with rest api.

And if u are tired of written, this:

```go

func SomeHandler(w http.ResponseWriter, r *http.Request) {

    w.Header().Add("Content-Type", "application/json")

    product := &product{"Smart TV", 50.00}
    
    bytes, err := json.Marshal(product)
    
    if err != nil {
    	w.WriteHeader(http.StatusInternalServerError)
        // this is super bad!
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
    "github.com/edermanoel94/rest-go"
    "net/http"
)

type product struct {
    Name  string `json:"name"`
    Price float32 `json:"price"`
}

func SomeHandler(w http.ResponseWriter, r *http.Request) {
    rest.Marshalled(w, &product{"Smart TV", 50.00}, http.StatusOK)
}
```

A payload send to your API and desarialize to a struct, to easy!

```go
package yours

import (
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
    err := rest.GetBody(r.Body, &product)
    if err != nil {
        // do stuff with error
    }
    // Do stuff...
}    
```

Working with [`mux`](https://github.com/gorilla/mux "API documentation") package to check if path variable exist.

```go
package yours

import (
    "github.com/edermanoel94/rest-go"
    "github.com/gorilla/mux"
    "net/http"
)

type product struct {
    Name  string `json:"name"`
    Price float32 `json:"price"`
}

// [GET] url: /product/{id}
func SomePostHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    err := rest.CheckPathVariables(params, "id")
    if err != nil {
        // do stuff with error
    }
}
```

TODO List
=========

- [ ] Working with pagination and test
- [x] Working with custom errors
- [ ] More tests
- [ ] Working with CheckPathVariables in Standard library

Installation
============

To install, use `go get`:

```
go get github.com/edermanoel94/rest-go
```

Example
=======

To install and build:

```
cd examples
```

And use dep for install packages:

```
dep ensure
```

Then, to run:

```
go run *.go || go build
```

Contributing
============

Please feel free to submit issues, fork the repository and send pull requests!

------

License
=======

This project is licensed under the terms of the MIT license.
