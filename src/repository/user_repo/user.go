package seller_repo

import (
	"context"
	u "rakamin-academy/src/entity/user"
)

type UserRepository interface {
	Create(ctx context.Context, seller u.User) (u.User, error)
	FindByEmail(ctx context.Context, email string) (u.User, error)
}
