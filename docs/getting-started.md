# Getting Started

## Install

```bash
go get github.com/fulldump/box
```

## Minimal server

```go
package main

import (
	"net/http"

	"github.com/fulldump/box"
)

func main() {
	b := box.NewBox()
	b.HandleFunc(http.MethodGet, "/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("ok"))
	})
	_ = b.ListenAndServe()
}
```

## Typed handlers

`box` can deserialize JSON request bodies into typed arguments and serialize your
return value into JSON response bodies.

```go
type In struct {
	Name string `json:"name"`
}

type Out struct {
	Message string `json:"message"`
}

b.Handle(http.MethodPost, "/hello", func(in In) Out {
	return Out{Message: "hello " + in.Name}
})
```

## Error flow

If your handler returns `error`, `box` stores it in context. Add interceptors such
as `box.PrettyError` to format and return errors consistently.

```go
b.Use(box.PrettyError)
b.Handle(http.MethodGet, "/fail", func() (any, error) {
	return nil, errors.New("something went wrong")
})
```
