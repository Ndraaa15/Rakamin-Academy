package seller_repo

import (
	"context"
	e "rakamin-academy/src/entity/seller"
)

type SellerRepository interface {
	Create(ctx context.Context, seller e.Seller) (e.Seller, error)
}
