package entity

import "time"

import vo "yatter-backend-go/app/domain/value-object"

type Status struct {
	ID        int       `json:"id,omitempty"`
	AccountID vo.AccountID `json:"account_id,omitempty" db:"account_id"`
	URL       *string   `json:"url,omitempty" db:"url"`
	Content   string    `json:"content,omitempty" db:"content"`
	CreateAt   time.Time `json:"create_at,omitempty" db:"create_at"`
}

func NewStatus(account_id vo.AccountID, content string) (*Status, error) {
	Status := &Status{
		AccountID: account_id,
		Content: content,
		CreateAt: time.Now(),
	}
	return Status, nil
}
