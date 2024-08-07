package usecase

import (
	"context"
	"log"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	Create(ctx context.Context, account_id int, content string) (*CreateStatusDTO, error)
}

type status struct {
	db          *sqlx.DB
	statusRepo repository.Status
}

type CreateStatusDTO struct {
	Status *object.Status
}

var _ Status = (*status)(nil)

func NewStatus(db *sqlx.DB, statusRepo repository.Status) *status {
	return &status{
		db:          db,
		statusRepo: statusRepo,
	}
}

func (s *status) Create(ctx context.Context, account_id int,content string) (*CreateStatusDTO, error) {
	st, err := object.NewStatus(account_id, content)
	if err != nil {
		return nil, err
	}

	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
        if p := recover(); p != nil {
            if rbErr := tx.Rollback(); rbErr != nil {
                log.Printf("rollback error: %v", rbErr)
            }
            panic(p)
        } else if err != nil {
            if rbErr := tx.Rollback(); rbErr != nil {
                log.Printf("rollback error: %v", rbErr)
            }
        } else {
            if commitErr := tx.Commit(); commitErr != nil {
                log.Printf("commit error: %v", commitErr)
            }
        }
    }()


	if err := s.statusRepo.Create(ctx, tx, st); err != nil {
		return nil, err
	}

	return &CreateStatusDTO{
		Status: st,
	}, nil
}

