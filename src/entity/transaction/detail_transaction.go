package transaction

import (
	m "rakamin-academy/src/entity/merchant"
	p "rakamin-academy/src/entity/product"
	"time"
)

type DetailTransaction struct {
	ID            uint           `json:"id" gorm:"type:int(10) UNSIGNED AUTO_INCREMENT;primaryKey;"`
	TransactionID uint           `json:"transaction_id" gorm:"type:int(10);NOT NULL;"`
	Transaction   Transaction    `json:"transaction" gorm:"foreignKey:TransactionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MerchantID    uint           `json:"merchant_id" gorm:"type:int(10);NOT NULL;"`
	Merchant      m.Merchant     `json:"merchant" gorm:"foreignKey:MerchantID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	LogProductID  uint           `json:"log_product_id" gorm:"type:int(10);NOT NULL;"`
	LogProduct    []p.LogProduct `json:"log_product" gorm:"foreignKey:LogProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Quantity      uint           `json:"quantity" gorm:"type:int(10);NOT NULL;"`
	TotalPrice    uint           `json:"total_price" gorm:"type:int(10);NOT NULL;"`
	UpdateAt      time.Time      `json:"update_at" gorm:"type:date;NOT NULL;DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;"`
	CreateAt      time.Time      `json:"create_at" gorm:"type:date;NOT NULL;DEFAULT:CURRENT_TIMESTAMP;"`
}
