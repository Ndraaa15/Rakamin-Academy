package seller_service

import (
	"context"
	e "rakamin-academy/src/entity/seller"
	"rakamin-academy/src/model"
)

type SellerService interface {
	Register(ctx context.Context, registerInput model.RegisterRequest) (e.Seller, error)
}
