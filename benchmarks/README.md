# Router Benchmarks

This module benchmarks `box` against popular Go HTTP routers:

- `gin`
- `chi`
- `echo`

## Run benchmarks

From repository root:

```bash
go test ./...
```

From this folder:

```bash
go test -run '^$' -bench '^BenchmarkRouters$' -benchmem
```

## Covered scenarios

- `StaticGET`: basic static route
- `PathParamGET`: route with one path parameter
- `JSONPOST`: JSON decode + business logic + JSON encode

See `router_benchmark_test.go` for exact route implementations.
