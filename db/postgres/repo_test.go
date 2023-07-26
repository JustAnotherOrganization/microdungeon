package postgres_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"justanother.org/microdungeon/db/postgres"
)

type fixture struct {
	pool *pgxpool.Pool
	log  *zap.Logger
	repo postgres.Repo
}

func newFixture(t *testing.T) fixture {
	t.Helper()

	var (
		f   fixture
		err error
	)

	f.log, err = zap.NewDevelopment()
	require.NoError(t, err)

	f.pool, err = pgxpool.New(context.Background(), postgres.DSN())
	require.NoError(t, err)

	t.Cleanup(f.pool.Close)

	f.repo = postgres.New(f.pool, f.log)

	return f
}
