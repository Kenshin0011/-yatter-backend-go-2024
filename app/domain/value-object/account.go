package valueobject

import "fmt"

type AccountID struct {
	id int
}

func NewAccountID(id int) (*AccountID, error) {
	if id < 1 {
		return nil, fmt.Errorf("ユーザーIDは1以上である必要があります: %d", id)
	}

	accountID  := &AccountID{
		id: id,
	}
	return accountID, nil
}

func (a AccountID) Value() int {
    return a.id
}

func (a AccountID) Equals(other AccountID) bool {
	return a.id == other.id
}

func (id *AccountID) Scan(value interface{}) error {
	switch v := value.(type) {
	case int64:
		*id = AccountID{id: int(v)}
	case int:
		*id = AccountID{id: v}
	default:
		return fmt.Errorf("unsupported type %T for AccountID", v)
	}
	return nil
}

func (a AccountID) String() string {
	return fmt.Sprintf("%d", a.id)
}

func (a AccountID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%d\"", a.id)), nil
}

func (a *AccountID) UnmarshalJSON(data []byte) error {
	var id int
	if _, err := fmt.Sscanf(string(data), "\"%d\"", &id); err != nil {
		return fmt.Errorf("AccountIDのUnmarshalに失敗しました: %w", err)
	}
	accountID, err := NewAccountID(id)
	if err != nil {
		return fmt.Errorf("AccountIDのUnmarshalに失敗しました: %w", err)
	}
	*a = *accountID
	return nil
}


