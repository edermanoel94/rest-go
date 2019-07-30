Rest GO - Helpful library  for Rest API
================================

Go code (golang) a package that provide many tools for working with rest api.

Get started:

  * Install Rest GO with [one line of code](#installation), or [update it with another](#staying-up-to-date)



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

    bytes, _ := json.Marshal(product)

    _, _ = rest.Json(w, bytes, http.StatusOK)
}
```

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
