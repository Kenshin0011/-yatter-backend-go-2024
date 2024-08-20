package repository

import (
	"context"

	"yatter-backend-go/app/domain/entity"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	// TODO: Add Other APIs
	Create(ctx context.Context, tx *sqlx.Tx, st *entity.Status) error
	FindByID(ctx context.Context, id string) (*entity.Status, error)
	FindPublicTimeline(ctx context.Context, limit int) ([]*entity.Status, error)
}
