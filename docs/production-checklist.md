# Production Checklist

Use this checklist before adopting `box` in critical services.

- Define request timeouts at server and reverse proxy layers.
- Add panic recovery middleware (`box.RecoverFromPanic`).
- Add consistent error formatting (`box.PrettyError`).
- Add access logging and request correlation IDs.
- Validate and sanitize all external input payloads.
- Enforce authentication and authorization with interceptors.
- Add readiness/liveness endpoints.
- Benchmark hot routes and set performance budgets.
- Publish and version your OpenAPI document.
- Configure CI with tests and race detector.

For WebSocket services:

- Set origin checks.
- Set read limits and deadlines.
- Handle connection lifecycle metrics.
