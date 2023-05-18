package transaction

import (
	s "rakamin-academy/src/entity/seller"
	"time"
)

type Transaction struct {
	ID          uint      `json:"id" gorm:"type:int(10) UNSIGNED AUTO_INCREMENT;primaryKey;"`
	SellerID    uint      `json:"seller_id" gorm:"type:int(10);NOT NULL;"`
	Seller      s.Seller  `json:"seller" gorm:"foreignKey:SellerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AddressID   uint      `json:"address_id" gorm:"NOT NULL;"`
	Address     Address   `json:"address" gorm:"foreignKey:AddressID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TotalPrice  uint      `json:"total_price" gorm:"NOTE NULL;"`
	InvoiceCode string    `json:"invoice_code" gorm:"type:varchar(255);NOT NULL;"`
	PaymentType string    `json:"payment_type" gorm:"type:varchar(255);NOT NULL;"`
	UpdateAt    time.Time `json:"update_at" gorm:"type:date;NOT NULL;DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;"`
	CreateAt    time.Time `json:"create_at" gorm:"type:date;NOT NULL;DEFAULT:CURRENT_TIMESTAMP;"`
}
