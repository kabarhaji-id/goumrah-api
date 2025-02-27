package database

import (
	"context"
)

type UnitOfWork struct {
	db DB
}

func NewUnitOfWork(db DB) *UnitOfWork {
	uow := new(UnitOfWork)
	uow.db = db

	return uow
}

func (u *UnitOfWork) Do(ctx context.Context, fn func(ctx context.Context, db DB) error) error {
	tx, err := u.db.Begin(ctx)
	if err != nil {
		return err
	}

	if err = fn(ctx, tx); err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return rollbackErr
		}

		return err
	}

	return tx.Commit(ctx)
}
