# Middleware

In `box`, middlewares are called interceptors.

## Global interceptors

```go
b.Use(box.AccessLog, box.PrettyError)
```

These run for every route.

## Group interceptors

```go
api := b.Group("/api")
api.Use(AuthInterceptor)
```

These run only for that group subtree.

## Route interceptors

```go
b.Handle(http.MethodPost, "/articles", CreateArticle).Use(ValidateArticle)
```

These run only for one route.

## Custom interceptor shape

```go
func Trace(next box.H) box.H {
	return func(ctx context.Context) {
		start := time.Now()
		next(ctx)
		log.Printf("request took=%s", time.Since(start))
	}
}
```

Interceptors can read and write request/response objects through helper functions
like `box.GetRequest(ctx)` and `box.GetResponse(ctx)`.
