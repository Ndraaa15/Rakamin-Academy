package seller_service

import (
	"context"
	"errors"
	"rakamin-academy/sdk/validator"
	e "rakamin-academy/src/entity/seller"
	"rakamin-academy/src/model"
	r "rakamin-academy/src/repository/seller_repo"
)

type Seller struct {
	sr r.SellerRepository
}

func NewSellerService(sellerRepo r.SellerRepository) SellerService {
	return &Seller{sellerRepo}
}

func (ss *Seller) Register(ctx context.Context, registerInput model.RegisterRequest) (e.Seller, error) {
	var seller e.Seller

	if !validator.EmailValidator(registerInput.Email) {
		return seller, errors.New("invalid format email")
	}

	if !validator.PasswordValidator(registerInput.Password) {
		return seller, errors.New("invalid format password")
	}

	if !validator.PhoneValidator(registerInput.Contact) {
		return seller, errors.New("invalid format phone number")
	}

	seller = e.Seller{
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

	seller, err := ss.sr.Create(ctx, seller)

	if err != nil {
		return seller, errors.New("failed to register")
	}

	return seller, nil
}
