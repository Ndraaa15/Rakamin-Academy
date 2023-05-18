package seller_service

import (
	r "rakamin-academy/src/repository/seller_repo"
)

type Seller struct {
	repository r.SellerRepository
}

func NewSellerService(sellerRepo r.SellerRepository) SellerService {
	return &Seller{sellerRepo}
}

func (seller *Seller) Register() {}
