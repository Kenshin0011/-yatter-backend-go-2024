package object

import "time"

type Status struct {
	ID        int       `json:"id,omitempty"`
	AccountID int       `json:"account_id,omitempty" db:"account_id"`
	URL       *string   `json:"url,omitempty" db:"url"`
	Content   string    `json:"status"`
	CreatAt   time.Time `json:"creat_at,omitempty" db:"creat_at"`
}

func NewStatus(account_id int, content string) (*Status, error) {
	Status := &Status{
		AccountID: account_id,
		Content: content,
		CreatAt: time.Now(),
	}
	return Status, nil
}
