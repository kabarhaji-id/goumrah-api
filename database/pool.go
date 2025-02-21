package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kabarhaji-id/goumrah-api/config"
)

var Pool *pgxpool.Pool

func InitPool(cfg config.Config) error {
	if Pool == nil {
		var err error
		Pool, err = pgxpool.New(context.Background(), cfg.PostgresDSN)
		if err != nil {
			return err
		}
	}

	return nil
}
