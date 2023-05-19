package seller_repo

import (
	"context"
	"rakamin-academy/database"
	u "rakamin-academy/src/entity/user"
)

type User struct {
	sql      *database.SQL
	supabase *database.Supabase
}

func NewUserRepository(sql *database.SQL, supabase *database.Supabase) UserRepository {
	return &User{sql, supabase}
}

func (ur *User) Create(ctx context.Context, seller u.User) (u.User, error) {
	if err := ur.sql.Debug().WithContext(ctx).Create(&seller).Error; err != nil {
		return seller, err
	}

	return seller, nil
}

func (ur *User) FindByEmail(ctx context.Context, email string) (u.User, error) {
	var seller u.User

	if err := ur.sql.Debug().WithContext(ctx).Where("email = ?", email).First(&seller).Error; err != nil {
		return seller, err
	}

	return seller, nil
}
