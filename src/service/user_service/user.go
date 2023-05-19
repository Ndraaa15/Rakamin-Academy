package seller_service

import (
	"context"
	u "rakamin-academy/src/entity/user"
	"rakamin-academy/src/model"
)

type UserService interface {
	Register(ctx context.Context, registerInput model.RegisterRequest) (u.User, error)
}
