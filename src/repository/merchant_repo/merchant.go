package merchant_repo

import (
	"context"
	m "rakamin-academy/src/entity/merchant"
)

type MerchantRepository interface {
	Create(ctx context.Context, merchant m.Merchant) (m.Merchant, error)
}
