package seller_service

import (
	"context"
	"errors"
	"rakamin-academy/sdk/password"
	"rakamin-academy/sdk/validator"
	m "rakamin-academy/src/entity/merchant"
	u "rakamin-academy/src/entity/user"
	"rakamin-academy/src/model"
	mr "rakamin-academy/src/repository/merchant_repo"
	ur "rakamin-academy/src/repository/user_repo"
)

type User struct {
	ur ur.UserRepository
	mr mr.MerchantRepository
}

func NewUserService(ur ur.UserRepository, mr mr.MerchantRepository) UserService {
	return &User{
		ur: ur,
		mr: mr,
	}
}

func (us *User) Register(ctx context.Context, registerInput model.RegisterRequest) (u.User, error) {

	var user u.User

	var merchant m.Merchant

	if !validator.EmailValidator(registerInput.Email) {
		return user, errors.New("INVALID FORMAT EMAIL")
	}

	if !validator.PasswordValidator(registerInput.Password) {
		return user, errors.New("INVALID FORMAT PASSWORD")
	}

	if !validator.PhoneValidator(registerInput.Contact) {
		return user, errors.New("INVALID FORMAT PHONE NUMBER")
	}

	password, err := password.HashPassword(registerInput.Password)

	if err != nil {
		return user, errors.New("FAILED TO HASH PASSWORD")
	}

	user = u.User{
		Name:       registerInput.Name,
		Email:      registerInput.Email,
		Contact:    registerInput.Contact,
		Password:   password,
		BirthDate:  registerInput.BirthDate,
		Gender:     registerInput.Gender,
		About:      registerInput.About,
		Job:        registerInput.Job,
		ProvinceID: registerInput.ProvinceID,
		CityID:     registerInput.CityID,
	}

	user, err = us.ur.Create(ctx, user)

	if err != nil {
		return user, errors.New("FAILED TO CREATE USER")
	}

	merchant = m.Merchant{
		UserID: user.ID,
		User:   user,
	}

	_, err = us.mr.Create(ctx, merchant)

	if err != nil {
		return user, errors.New("FAILED TO CREATE MERCHANT")
	}

	return user, nil

}

func (us *User) Login(ctx context.Context, loginInput model.LoginRequest) (model.LoginResponse, error) {
	var loginResponse model.LoginResponse

	if !validator.EmailValidator(loginInput.Email) {
		return loginResponse, errors.New("invalid format email")
	}

	if !validator.PasswordValidator(loginInput.Password) {
		return loginResponse, errors.New("invalid format password")
	}

	user, err := us.ur.FindByEmail(ctx, loginInput.Email)

	if err != nil {
		return loginResponse, errors.New("user not found")
	}

	loginResponse = model.LoginResponse{
		User:  user,
		Token: "token",
	}

	return loginResponse, nil
}
