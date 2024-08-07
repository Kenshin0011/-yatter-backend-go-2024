package object

import "time"

type Status struct {
	ID        int       `json:"id,omitempty"`
	AccountID int       `json:"account_id,omitempty" db:"account_id"`
	URL       *string   `json:"url,omitempty" db:"url"`
	Content   string    `json:"contenr,omitempty" db:"content"`
	CreateAt   time.Time `json:"create_at,omitempty" db:"create_at"`
}

func NewStatus(account_id int, content string) (*Status, error) {
	Status := &Status{
		AccountID: account_id,
		Content: content,
		CreateAt: time.Now(),
	}
	return Status, nil
}
