package auth

import (
	"context"
	"yatter-backend-go/app/domain/entity"
)

var ContextKey = new(struct{})

// Read Account data from authorized request
func AccountOf(ctx context.Context) *entity.Account {
	if cv := ctx.Value(ContextKey); cv == nil {
		return nil

	} else if account, ok := cv.(*entity.Account); !ok {
		return nil

	} else {
		return account

	}
}
