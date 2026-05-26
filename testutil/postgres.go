package testutil

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

func NewTestPgxPool(t *testing.T) *pgxpool.Pool {
	t.Helper()

	postgresContainer, err := postgres.Run(t.Context(),
		"postgres:17-alpine",
		postgres.WithDatabase("laughing_goggles_test"),
		postgres.BasicWaitStrategies(),
	)
	require.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, testcontainers.TerminateContainer(postgresContainer))
	})

	connectionString, err := postgresContainer.ConnectionString(t.Context())
	require.NoError(t, err)

	pool, err := pgxpool.New(t.Context(), connectionString)
	require.NoError(t, err)
	t.Cleanup(pool.Close)

	_, thisFile, _, _ := runtime.Caller(0)
	migrationsDir := filepath.Join(filepath.Dir(thisFile), "../db/migrations")

	goose.SetLogger(goose.NopLogger())

	db := stdlib.OpenDBFromPool(pool)
	err = goose.Up(db, migrationsDir)
	require.NoError(t, err)

	return pool
}
