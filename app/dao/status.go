package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type status struct {
	db *sqlx.DB
}

// Create accout repository
func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

func (s *status) Create(ctx context.Context, tx *sqlx.Tx, st *object.Status) error {
	q := `
	INSERT INTO status (account_id, url, content, create_at)
	VALUES (:account_id, :url, :content, :create_at);
	`

	q, params, err := sqlx.Named(q, map[string]interface{}{
		"account_id": st.AccountID,
		"url": st.URL,
		"content": st.Content,
		"create_at": st.CreateAt,
	})

	if err != nil {
		return fmt.Errorf("failed to insert status: %w", err)
	}
		
	_, err = tx.ExecContext(ctx, q, params)
	if err != nil {
		return fmt.Errorf("%s: %v",q, err)
	}

	return nil
}

func (s *status) FindByID(ctx context.Context, id string) (*object.Status, error) {
	entity := new(object.Status)
	err := s.db.QueryRowxContext(ctx, "select * from status where id = ?", id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no status found with the given id: %w", err)
		}

		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}
	return entity, nil
}

func (s *status) FindPublicTimeline(ctx context.Context, limit int) ([]*object.Status, error) {
	var entities []*object.Status

	err := s.db.SelectContext(ctx, &entities, "select * from status order by id desc limit ?", limit)
	if err != nil {
		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}
	return entities, nil
}