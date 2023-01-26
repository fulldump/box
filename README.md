# Box

<p align="center">
<a href="https://app.travis-ci.com/fulldump/box" rel="nofollow"><img src="https://app.travis-ci.com/fulldump/box.svg?branch=master" alt="Build Status"></a>
<a href="https://goreportcard.com/report/github.com/fulldump/box"><img src="https://goreportcard.com/badge/github.com/fulldump/box"></a>
<a href="https://godoc.org/github.com/fulldump/box"><img src="https://godoc.org/github.com/fulldump/box?status.svg" alt="GoDoc"></a>
</p>



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