# Contributing to PuzzleSmith

Everyone is welcome to contribute! Help improve the Go processing pipeline, backend API, frontend, performance, CI/CD, tests, or documentation. Bug reports, ideas, and recommendations are welcome too.

## Getting started

1. Fork and clone the repository, then create a branch for your change.
2. Follow the [README](README.md#local-development) for local setup.
3. Make a focused change, follow existing conventions, and format changed Go files with `gofmt`.
4. Run the relevant checks below and open a pull request targeting `main`. Explain your change and what you verified; include screenshots for UI changes when useful.

## Checks

Backend, from `apps/api`:

```sh
go vet ./...
go test ./...
go build -o bin/api .
```

Frontend, from `apps/web`:

```sh
pnpm install --frozen-lockfile
pnpm lint
pnpm build
```

Keep secrets and local `.env` files out of commits. For bug reports, include reproduction steps and expected versus actual behavior, with private data removed.
