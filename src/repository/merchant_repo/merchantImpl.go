package merchant_repo

import (
	"context"
	"rakamin-academy/database"
	m "rakamin-academy/src/entity/merchant"
)

type Merchant struct {
	sql      *database.SQL
	supabase *database.Supabase
}

func NewMerchantRepository(sql *database.SQL, supabase *database.Supabase) MerchantRepository {
	return &Merchant{sql, supabase}
}

func (mr *Merchant) Create(ctx context.Context, merchant m.Merchant) (m.Merchant, error) {

	if err := mr.sql.Debug().WithContext(ctx).Create(&merchant).Error; err != nil {
		return merchant, err
	}

	return merchant, nil

}
