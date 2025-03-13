package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDatabase(dsn string) (ClosableDB, error) {
	return pgxpool.New(context.Background(), dsn)
}
