# OpenAPI

You can generate OpenAPI schemas from route and type definitions using
`github.com/fulldump/box/boxopenapi`.

## Basic usage

```go
import "github.com/fulldump/box/boxopenapi"

spec := boxopenapi.Spec(b)
spec.Info.Title = "My API"
spec.Info.Version = "1.0.0"

b.Handle(http.MethodGet, "/openapi.json", func() any {
	return spec
})
```

## Good practices

- Add a service-level interceptor to standardize error responses.
- Keep request/response DTOs explicit and versioned.
- Expose `/openapi.json` and a UI (Swagger UI or Redoc) in non-private APIs.
