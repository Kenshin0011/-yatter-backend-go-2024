package usecase

import (
	"context"
	"fmt"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Account interface {
	Create(ctx context.Context, username, password string) (*CreateAccountDTO, error)
	FindByUsername(ctx context.Context, username string) (*GetAccountDTO, error)
}

type account struct {
	db          *sqlx.DB
	accountRepo repository.Account
}

type CreateAccountDTO struct {
	Account *object.Account
}

type GetAccountDTO struct {
	Account *object.Account
}

var _ Account = (*account)(nil)

func NewAcocunt(db *sqlx.DB, accountRepo repository.Account) *account {
	return &account{
		db:          db,
		accountRepo: accountRepo,
	}
}

func (a *account) Create(ctx context.Context, username, password string) (*CreateAccountDTO, error) {
	acc, err := object.NewAccount(username, password)
	if err != nil {
		return nil, err
	}

	unitOfWork := NewUnitOfWork(a.db)
	err = unitOfWork.Do(ctx, func(tx *sqlx.Tx) error {
        err = a.accountRepo.Create(ctx, tx, acc)
		fmt.Println(err)
        if err != nil {
            return fmt.Errorf("failed to create account: %w", err)
        }
        return nil
    })

    if err != nil {
        return nil, err
    }

	return &CreateAccountDTO{
		Account: acc,
	}, nil
}

func (a *account) FindByUsername(ctx context.Context, username string) (*GetAccountDTO, error) {
	acc, err := a.accountRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return &GetAccountDTO{
		Account: acc,
	}, nil
}
