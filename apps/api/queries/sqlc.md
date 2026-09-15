# sqlc + goose workflow

This repo uses PostgreSQL with `goose` for migrations and `sqlc` for type-safe queries.

## 1) Create a migration
`goose -dir db/migrations create <name> sql`

Edit the new file and add your schema changes in the `Up` section and rollback in `Down`.

## 2) Apply migrations
`goose -dir db/migrations postgres "$DATABASE_URL" up`

## 3) Write queries
Put SQL in `queries/query.sql` using `sqlc` query annotations, for example:
- `-- name: GetThing :one`
- `-- name: ListThings :many`

## 4) Generate code
`sqlc generate`

## 5) Repeat
Schema change = new migration + `goose up`.
Query change = edit query file + `sqlc generate`.
