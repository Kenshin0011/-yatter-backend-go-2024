package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/entity"
	"yatter-backend-go/app/domain/repository"
	vo "yatter-backend-go/app/domain/value-object"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	account struct {
		db *sqlx.DB
	}
)

var _ repository.Account = (*account)(nil)

// Create accout repository
func NewAccount(db *sqlx.DB) *account {
	return &account{db: db}
}

func (a *account) Create(ctx context.Context, tx *sqlx.Tx, acc *entity.Account) error {
	result, err := tx.Exec("insert into account (username, password_hash, display_name, avatar, header, note, create_at) values (?, ?, ?, ?, ?, ?, ?)",
		acc.Username, acc.PasswordHash, acc.DisplayName, acc.Avatar, acc.Header, acc.Note, acc.CreateAt)
	if err != nil {
		return fmt.Errorf("failed to insert account: %w", err)
	}

	id, err := result.LastInsertId()
    if err != nil {
        return fmt.Errorf("挿入IDの取得に失敗しました: %w", err)
    }

    accountID, err := vo.NewAccountID(int(id))
    if err != nil {
        return fmt.Errorf("AccountIDの生成に失敗しました: %w", err)
    }

	acc.ID = *accountID
	return nil
}

func (a *account) FindByUsername(ctx context.Context, username string) (*entity.Account, error) {
	entity := new(entity.Account)
	err := a.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no user found with the given username: %w", err)
		}

		return nil, fmt.Errorf("failed to find account from db: %w", err)
	}

	return entity, nil
}
