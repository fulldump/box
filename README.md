# box

<p align="center">
  <img src="logo.png" alt="box logo" width="220" />
</p>

<p align="center">
  <a href="https://pkg.go.dev/github.com/fulldump/box"><img src="https://pkg.go.dev/badge/github.com/fulldump/box.svg" alt="Go Reference"></a>
  <a href="https://goreportcard.com/report/github.com/fulldump/box"><img src="https://goreportcard.com/badge/github.com/fulldump/box" alt="Go Report Card"></a>
  <a href="https://github.com/fulldump/box/actions/workflows/ci.yml"><img src="https://github.com/fulldump/box/actions/workflows/ci.yml/badge.svg" alt="CI"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="MIT License"></a>
</p>

`box` is a fast, typed HTTP router for Go with clean middleware composition,
automatic JSON handling, and built-in OpenAPI generation.

It is already battle-tested in production-like environments and now aims to be
the most ergonomic open-source router for teams that care about productivity
without losing control.

## Why box

- Typed handlers: receive typed request bodies and return typed responses.
- Familiar API: still works with standard `http.HandlerFunc` handlers.
- Composable middleware: global, group, and route-level interceptors.
- URL params and route groups: ergonomic API design for real services.
- Built-in OpenAPI generation: publish docs from existing route definitions.
- WebSocket support: native route helper based on the same router tree.

## Install

```bash
go get github.com/fulldump/box
```

## Quickstart

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

	_ = b.ListenAndServe() // http://localhost:8080
}
```

## Typed JSON handlers

```go
type CreateArticleRequest struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type Article struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

b := box.NewBox()

b.Handle(http.MethodPost, "/articles", func(input CreateArticleRequest) (Article, error) {
	return Article{
		ID:    "a-1",
		Title: input.Title,
		Text:  input.Text,
	}, nil
})
```

## URL parameters

```go
b := box.NewBox()

b.Handle(http.MethodGet, "/articles/{article-id}", func(r *http.Request) string {
	articleID := r.PathValue("article-id") // Go >= 1.22
	return "article " + articleID
})

// Go < 1.22
// articleID := box.Param(r, "article-id")
```

## Interceptors (middlewares)

```go
b := box.NewBox()

b.Use(box.AccessLog, box.PrettyError)

api := b.Group("/api")
api.Use(box.SetResponseHeader("X-Service", "articles"))

api.Handle(http.MethodGet, "/articles", func() []string {
	return []string{"a-1", "a-2"}
})
```

## WebSocket routes

```go
import (
	"context"

	"github.com/gorilla/websocket"
)

b := box.NewBox()

b.HandleWebSocket("/ws/echo", func(_ context.Context, conn *websocket.Conn) {
	for {
		messageType, payload, err := conn.ReadMessage()
		if err != nil {
			return
		}
		_ = conn.WriteMessage(messageType, payload)
	}
})
```

You can also provide custom `box.WebSocketOptions` with your own `websocket.Upgrader`
and error strategy.

## OpenAPI generation

```go
import "github.com/fulldump/box/boxopenapi"

spec := boxopenapi.Spec(b)
spec.Info.Title = "Articles API"
spec.Info.Version = "1.0.0"

b.Handle(http.MethodGet, "/openapi.json", func() any {
	return spec
})
```

## Documentation

- Project docs: `docs/`
- Publishable website: `website/`
- Benchmark suite: `benchmarks/`
- Examples in tests: `example_*_test.go`, `examples_*_test.go`

## Benchmarks

Run the comparative benchmark suite against `box`, `gin`, `chi`, and `echo`:

```bash
cd benchmarks
go test -run '^$' -bench '^BenchmarkRouters$' -benchmem
```

Latest baseline and methodology are published in `docs/benchmarks.md`.

## Stability and compatibility

- Go module: `github.com/fulldump/box`
- Current minimum Go version: `1.19`
- Follows semantic versioning for public APIs

## Contributing

Contributions are welcome. Please read:

- `CONTRIBUTING.md`
- `CODE_OF_CONDUCT.md`
- `SECURITY.md`

## Roadmap

- Close benchmark performance gap with gin/chi/echo while preserving API ergonomics.
- More first-party interceptors (auth, rate-limiting, observability).
- Improved OpenAPI customization hooks.
- Extra guides and production deployment recipes.

## License

MIT. See `LICENSE`.
