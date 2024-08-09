package usecase

import (
	"context"
	"testing"
	"yatter-backend-go/app/domain/object"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

type MockStatusRepo struct {
	mockCreateFunc    func(ctx context.Context, tx *sqlx.Tx, st *object.Status) error
	mockFindByIDFunc  func(ctx context.Context, id string) (*object.Status, error)
	mockPublicTimelineFunc func(ctx context.Context, limit int) ([]*object.Status, error)
}

func (m *MockStatusRepo) FindPublicTimeline(ctx context.Context, limit int) ([]*object.Status, error) {
	if m.mockPublicTimelineFunc != nil {
		return m.mockPublicTimelineFunc(ctx, limit)
	}
	return nil, nil
}

func (m *MockStatusRepo) FindByID(ctx context.Context, id string) (*object.Status, error) {
	if m.mockPublicTimelineFunc != nil {
		return m.mockFindByIDFunc(ctx, id)
	}
	return nil, nil
}

func (m *MockStatusRepo) Create(ctx context.Context, tx *sqlx.Tx, st *object.Status) error {
	if m.mockCreateFunc != nil {
		return m.mockCreateFunc(ctx, tx, st)
	}
	return nil
}

type MockUnitOfWork struct {
	mockDoFunc func(ctx context.Context, f func(tx *sqlx.Tx) error) error
}

func (m *MockUnitOfWork) Do(ctx context.Context, f func(tx *sqlx.Tx) error) error {
	if m.mockDoFunc != nil {
		return m.mockDoFunc(ctx, f)
	}
	return nil
}

func TestStatusUsecase_Create(t *testing.T) {
	ctx := context.Background()

	t.Run("正常系: Status情報を正常に返すこと", func(t *testing.T) {
		account_id := 1
		content := "test content"

		mockDB := sqlx.NewDb(nil, "sqlmock")
		mockStatusRepo := MockStatusRepo{
			mockCreateFunc: func(ctx context.Context, tx *sqlx.Tx, st *object.Status) error {
				return nil
			},
		}
		mockUnitOfWork := MockUnitOfWork{
			mockDoFunc: func(ctx context.Context, f func(tx *sqlx.Tx) error) error {
				return nil
			},
		}

		sut := NewStatus(mockDB, &mockStatusRepo, &mockUnitOfWork)

		got, err := sut.Create(ctx, account_id, content)
		assert.NoError(t, err)
		want, err := object.NewStatus(account_id, content)
		assert.Equal(t, want.AccountID, got.Status.AccountID)
		assert.Equal(t, want.Content, got.Status.Content)
	})

	t.Run("異常系: StatusRepository.Create()でエラーが発生した場合、エラーを返すこと", func(t *testing.T) {

	})
}