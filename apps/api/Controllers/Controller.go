package Controllers

import (
	"github.com/RajanDhamala/puzzleSmith/internal/db"

	"github.com/RajanDhamala/go-stockfish"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Controller struct {
	queries   *db.Queries
	pool      *pgxpool.Pool
	stockfish *stockfish.Client
}

func NewController(queries *db.Queries, pool *pgxpool.Pool, client *stockfish.Client) *Controller {
	return &Controller{queries: queries, pool: pool, stockfish: client}
}
