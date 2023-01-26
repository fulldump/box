# Box

[![Build Status](https://app.travis-ci.com/fulldump/box.svg?branch=master)](https://app.travis-ci.com/fulldump/box)

Box is an HTTP router to speed up development. Box supports URL parameters, interceptors, magic handlers
and introspection documentation.

## Getting started

```go
package main

import (
	"github.com/fulldump/box"
)

func main() {
	
	b := box.NewBox()
	b.HandleFunc("GET", "/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("World!"))
	})
	b.ListenAndServe() // listening at http://localhost:8080

}
```

## Getting started

```go
package main

import (
	"github.com/fulldump/box"
)

func main() {
	
	b := box.NewBox()
	b.HandleFunc("GET", "/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("World!"))
	})
	b.ListenAndServe() // listening at http://localhost:8080

}
```

## Sending JSON

```go
package main

import (
	"github.com/fulldump/box"
)

type MyResponse struct {
	Name string
	Age  int
}

func main() {

	b := box.NewBox()
	b.Handle("GET", "/hello", func(w http.ResponseWriter, r *http.Request) MyResponse {
		return MyResponse{
			Name: "Fulanez",
			Age:  33,
		}
	})
	b.ListenAndServe()

}
```