package seller_repo

import (
	"context"
	"rakamin-academy/database"
	e "rakamin-academy/src/entity/seller"
)

type Seller struct {
	sql      *database.SQL
	supabase *database.Supabase
}

func NewSellerRepository(sql *database.SQL, supabase *database.Supabase) SellerRepository {
	return &Seller{sql, supabase}
}

func (sr *Seller) Create(ctx context.Context, seller e.Seller) (e.Seller, error) {
	if err := sr.sql.Debug().WithContext(ctx).Create(&seller).Error; err != nil {
		return seller, err
	}

	return seller, nil
}
