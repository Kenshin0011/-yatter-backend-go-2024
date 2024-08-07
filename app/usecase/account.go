package usecase

import (
	"context"
	"log"
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

	tx, err := a.db.Beginx()
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

	if err := a.accountRepo.Create(ctx, tx, acc); err != nil {
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
