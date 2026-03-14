# Routing

## Bind routes

Use `Handle` for introspected handlers and `HandleFunc` for standard
`http.HandlerFunc`.

```go
b.Handle(http.MethodGet, "/articles", ListArticles)
b.HandleFunc(http.MethodGet, "/health", Health)
```

## URL parameters

Define placeholders with `{name}`.

```go
b.Handle(http.MethodGet, "/articles/{article-id}", func(r *http.Request) string {
	return r.PathValue("article-id")
})
```

For Go versions before `1.22`, use:

```go
id := box.Param(r, "article-id")
```

## Groups

Group routes by prefix and apply interceptors per group.

```go
api := b.Group("/api")
v1 := api.Group("/v1")

v1.Handle(http.MethodGet, "/articles", ListArticles)
```

## Method matching

`box` matches by path and method. If path exists but method does not, it triggers
the configured method-not-allowed handler.
