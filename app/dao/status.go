package dao

import (
	"context"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	status struct {
		db *sqlx.DB
	}
)

var _ repository.Status = (*status)(nil)

// Create accout repository
func NewStatus(db *sqlx.DB) *status {
	return &status{db: db}
}

func (a *status) Create(ctx context.Context, tx *sqlx.Tx, st *object.Status) error {
	_, err := a.db.Exec("insert into status (account_id, content, create_at) values (?, ?, ?)", st.AccountID, st.Content, st.CreatAt)
	if err != nil {
		return fmt.Errorf("failed to insert status: %w", err)
	}

	return nil
}
