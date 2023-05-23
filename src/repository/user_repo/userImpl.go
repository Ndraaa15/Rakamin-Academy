package seller_repo

import (
	"context"
	"errors"
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

func (ur *User) Create(ctx context.Context, user u.User) (u.User, error) {

	if err := ur.sql.Debug().WithContext(ctx).Create(&user).Error; err != nil {
		return user, errors.New("GORM ERROR : " + err.Error())
	}

	return user, nil

}

func (ur *User) FindByEmail(ctx context.Context, email string) (u.User, error) {

	var user u.User

	if err := ur.sql.Debug().WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return user, errors.New("GORM ERROR : " + err.Error())
	}

	return user, nil

}

func (ur *User) FindByID(ctx context.Context, userID uint) (u.User, error) {

	var user u.User

	if err := ur.sql.Debug().WithContext(ctx).Where("id = ?", userID).First(&user).Error; err != nil {
		return user, errors.New("GORM ERROR : " + err.Error())
	}

	return user, nil

}
