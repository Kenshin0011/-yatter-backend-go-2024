package usecase

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	Create(ctx context.Context, account_id int, content string) (*StatusDTO, error)
	FindByID(ctx context.Context, id string) (*StatusDTO, error)
	FindPublicTimeline(ctx context.Context, limit int) ([]*object.Status, error)
}

type StatusDTO struct {
	Status object.Status
}

type status struct {
	statusRepo 	repository.Status
	unitOfWork  UnitOfWork
}

func NewStatus(db *sqlx.DB, statusRepo repository.Status, unitOfWork UnitOfWork) Status {
	return &status{statusRepo, unitOfWork}
}

func (s *status) Create(ctx context.Context, account_id int, content string) (*StatusDTO, error) {
	st, err := object.NewStatus(account_id, content)
	if err != nil {
		return nil, err
	}

	// if err := s.statusRepo.Create(ctx, nil, st); err != nil {
	// 	return nil, fmt.Errorf("failed to create status:(accountID: %d): %v", st.AccountID ,err)
	// }

	err = s.unitOfWork.Do(ctx, func(tx *sqlx.Tx) error {
		err = s.statusRepo.Create(ctx, tx, st)
		if err != nil {
			return err
		}
		return nil
	})

	return &StatusDTO{
		Status: *st,
	}, nil
}

func (s *status) FindByID(ctx context.Context, id string) (*StatusDTO, error) {
	st, err := s.statusRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &StatusDTO{
		Status: *st,
	}, nil
}

func (s *status) FindPublicTimeline(ctx context.Context, limit int) ([]*object.Status, error) {
	st, err := s.statusRepo.FindPublicTimeline(ctx, limit)
	if err != nil {
		return nil, err
	}

	return st, nil
}
