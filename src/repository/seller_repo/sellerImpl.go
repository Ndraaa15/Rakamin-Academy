package seller_repo

import "rakamin-academy/database"

type Seller struct {
	sql      *database.SQL
	supabase *database.Supabase
}

func NewSellerRepository(sql *database.SQL, supabase *database.Supabase) SellerRepository {
	return &Seller{sql, supabase}
}
