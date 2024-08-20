package repository

import (
	"context"

	"yatter-backend-go/app/domain/entity"

	"github.com/jmoiron/sqlx"
)

type Account interface {
	// TODO: Add Other APIs
	Create(ctx context.Context, tx *sqlx.Tx, acc *entity.Account) error
	FindByUsername(ctx context.Context, username string) (*entity.Account, error)
}
