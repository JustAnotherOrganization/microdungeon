package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

const (
	byIdClause   = "WHERE id = $1"
	byNameClause = "WHERE name = $1"
)

type Repo struct {
	pool *pgxpool.Pool
	log  *zap.Logger
}

func New(pool *pgxpool.Pool, log *zap.Logger) Repo {
	return Repo{
		pool: pool,
		log:  log,
	}
}
