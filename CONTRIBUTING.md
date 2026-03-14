# Contributing to box

Thanks for your interest in contributing.

## Development setup

1. Install Go `1.19+`.
2. Clone the repository.
3. Run tests:

```bash
go test ./...
```

## Pull request guidelines

- Keep PRs focused and small when possible.
- Add tests for behavior changes.
- Update docs (`README.md` and `docs/`) when adding user-facing features.
- Prefer backward-compatible API changes.

## Commit conventions

Use clear messages that explain the intent, for example:

- `feat: add websocket route helper`
- `docs: improve getting started guide`
- `fix: preserve response headers in serializer`

## Reporting issues

When opening an issue, include:

- Go version
- Operating system
- Minimal reproducible example
- Expected behavior vs actual behavior
