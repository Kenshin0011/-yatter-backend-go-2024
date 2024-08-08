package usecase

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UnitOfWork interface {
	Do(ctx context.Context, f func(tx *sqlx.Tx) error) error
}

type unitOfWork struct {
	db *sqlx.DB
}

func NewUnitOfWork(db *sqlx.DB) UnitOfWork {
	return &unitOfWork{
		db: db,
	}
}

func (u *unitOfWork) Do(ctx context.Context, f func(tx *sqlx.Tx) error) error {
	tx, err := u.db.Beginx()
	if err != nil {
		return err
	}

	err = f(tx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("rollback error: %w", rbErr)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}