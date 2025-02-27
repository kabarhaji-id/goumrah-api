package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type resultPostgres struct {
	commandTag pgconn.CommandTag
}

func (t *resultPostgres) LastInsertId() (int64, error) {
	return 0, nil
}

func (t *resultPostgres) RowsAffected() (int64, error) {
	return t.commandTag.RowsAffected(), nil
}

type rowsPostgres struct {
	rows pgx.Rows
}

func (r *rowsPostgres) Close() error {
	r.rows.Close()

	return nil
}

func (r *rowsPostgres) Err() error {
	return r.rows.Err()
}

func (r *rowsPostgres) Next() bool {
	return r.rows.Next()
}

func (r *rowsPostgres) Scan(dest ...any) error {
	return r.rows.Scan(dest...)
}

type txPostgres struct {
	tx pgx.Tx
}

func (t *txPostgres) Exec(ctx context.Context, query string, args ...interface{}) (Result, error) {
	result := new(resultPostgres)

	var err error
	result.commandTag, err = t.tx.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *txPostgres) Query(ctx context.Context, query string, args ...interface{}) (Rows, error) {
	rows := new(rowsPostgres)

	var err error
	rows.rows, err = t.tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (t *txPostgres) QueryRow(ctx context.Context, query string, args ...interface{}) Row {
	return t.tx.QueryRow(ctx, query, args...)
}

func (t *txPostgres) Begin(ctx context.Context) (Tx, error) {
	tx := new(txPostgres)

	var err error
	tx.tx, err = t.tx.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (t *txPostgres) Close(ctx context.Context) error {
	t.tx.Rollback(ctx)

	return nil
}

func (t *txPostgres) Commit(ctx context.Context) error {
	return t.tx.Commit(ctx)
}

func (t *txPostgres) Rollback(ctx context.Context) error {
	return t.tx.Rollback(ctx)
}

type dbPostgres struct {
	pool *pgxpool.Pool
}

func NewPostgres(dsn string) (DB, error) {
	db := new(dbPostgres)
	var err error
	db.pool, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *dbPostgres) Exec(ctx context.Context, query string, args ...interface{}) (Result, error) {
	result := new(resultPostgres)

	var err error
	result.commandTag, err = d.pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d *dbPostgres) Query(ctx context.Context, query string, args ...interface{}) (Rows, error) {
	rows := new(rowsPostgres)

	var err error
	rows.rows, err = d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (d *dbPostgres) QueryRow(ctx context.Context, query string, args ...interface{}) Row {
	return d.pool.QueryRow(ctx, query, args...)
}

func (d *dbPostgres) Begin(ctx context.Context) (Tx, error) {
	tx := new(txPostgres)

	var err error
	tx.tx, err = d.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (d *dbPostgres) Close(ctx context.Context) error {
	d.pool.Close()

	return nil
}
