package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
}

func Connect(ctx context.Context, databaseDSN string) (*pgxpool.Pool, error) {

	conn, err := pgxpool.New(ctx, databaseDSN)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}
