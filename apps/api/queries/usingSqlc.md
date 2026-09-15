# Using sqlc in this repo

This project generates type-safe query code with `sqlc`. The generated package is `internal/db` and the Go package name is `db` (from `sqlc.yaml`).

## 1) Generate code
Run from `go` directory:
`sqlc generate`

If you run from repo root:
`sqlc generate -f go/sqlc.yaml`

If you run from `go/queries`:
`sqlc generate -f ../sqlc.yaml`

## 2) Import and create queries
Import the generated package:
`"chess/internal/db"`

Create a `db.Queries` instance from a `pgxpool.Pool` or `pgx.Conn`:
`queries := db.New(pool)`

## 3) Call queries
Generated functions match the query names in `queries/query.sql`.
Examples based on current queries:
- `RegisterUser(ctx, db.RegisterUserParams{Fullname: ..., Email: ..., Password: ...})`
- `CheckIfusrExists(ctx, email)`
- `LoginUser(ctx, email)`

## 4) Error handling
Check errors from every call. Common cases:
- `pgx.ErrNoRows` when a `:one` query returns no rows.
- Unique constraint violations on insert.
- Connection errors if the pool is not initialized or closed.

## 5) Regenerate after changes
If you change SQL or migrations, run `sqlc generate` again.
