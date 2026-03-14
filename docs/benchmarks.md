# Benchmarks

`box` includes a public benchmark suite to compare against `gin`, `chi`, and
`echo` with reproducible code.

## How to run

```bash
cd benchmarks
go test -run '^$' -bench '^BenchmarkRouters$' -benchmem
```

## Methodology

- In-memory benchmark (`httptest`) to isolate router and handler overhead.
- Same high-level behavior for all frameworks in each scenario.
- Scenarios: static GET, path parameter GET, JSON POST.

## Baseline snapshot (2026-03-14)

Environment:

- OS: Linux
- CPU: 11th Gen Intel(R) Core(TM) i7-11800H @ 2.30GHz

Results (`ns/op`, lower is better):

| Scenario | box | gin | chi | echo |
| --- | ---: | ---: | ---: | ---: |
| StaticGET | 4924 | 2629 | 2990 | 2808 |
| PathParamGET | 4768 | 2706 | 3491 | 2646 |
| JSONPOST | 5207 | 4655 | 4467 | 4303 |

Results (`allocs/op`, lower is better):

| Scenario | box | gin | chi | echo |
| --- | ---: | ---: | ---: | ---: |
| StaticGET | 30 | 18 | 21 | 19 |
| PathParamGET | 34 | 18 | 23 | 19 |
| JSONPOST | 42 | 32 | 33 | 32 |

## Notes

- These numbers are a baseline and can vary by machine and Go version.
- The benchmark suite is intentionally checked into the repository for
  transparency and iterative optimization.
- The goal is to close the performance gap while preserving `box` ergonomics.
