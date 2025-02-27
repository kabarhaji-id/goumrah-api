package database

import (
	"context"
)

type DB interface {
	Exec(ctx context.Context, query string, args ...any) (Result, error)
	Query(ctx context.Context, query string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) Row
	Begin(ctx context.Context) (Tx, error)
	Close(ctx context.Context) error
}

type Tx interface {
	DB
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Rows interface {
	Close() error
	Err() error
	Next() bool
	Scan(dest ...any) error
}

type Row interface {
	Scan(dest ...any) error
}
