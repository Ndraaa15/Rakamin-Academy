package seller_service

import (
	"context"
	"errors"
	"rakamin-academy/sdk/validator"
	u "rakamin-academy/src/entity/user"
	"rakamin-academy/src/model"
	r "rakamin-academy/src/repository/user_repo"
)

type User struct {
	ur r.UserRepository
}

func NewUserService(userRepo r.UserRepository) UserService {
	return &User{userRepo}
}

func (us *User) Register(ctx context.Context, registerInput model.RegisterRequest) (u.User, error) {
	var seller u.User

	if !validator.EmailValidator(registerInput.Email) {
		return seller, errors.New("invalid format email")
	}

	if !validator.PasswordValidator(registerInput.Password) {
		return seller, errors.New("invalid format password")
	}

	if !validator.PhoneValidator(registerInput.Contact) {
		return seller, errors.New("invalid format phone number")
	}

	seller = u.User{
		Name:       registerInput.Name,
		Email:      registerInput.Email,
		Contact:    registerInput.Contact,
		Password:   registerInput.Password,
		BirthDate:  registerInput.BirthDate,
		Gender:     registerInput.Gender,
		About:      registerInput.About,
		Job:        registerInput.Job,
		ProvinceID: registerInput.ProvinceID,
		CityID:     registerInput.CityID,
		IsAdmin:    registerInput.IsAdmin,
	}

	seller, err := us.ur.Create(ctx, seller)

	if err != nil {
		return seller, errors.New("failed to register")
	}

	return seller, nil
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
