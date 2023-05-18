package product

import "time"

type LogProduct struct {
	ID            uint      `json:"id" gorm:"type:int(10) UNSIGNED AUTO_INCREMENT;primaryKey;"`
	ProductID     uint      `json:"product_id" gorm:"type:int(10);NOT NULL;"`
	Name          string    `json:"name" gorm:"type:varchar(255);NOT NULL;"`
	Slug          string    `json:"slug" gorm:"type:varchar(255);NOT NULL;"`
	ResellerPrice string    `json:"reseller_price" gorm:"type:varchar(255);NOT NULL;"`
	ConsumerPrice string    `json:"consumer_price" gorm:"type:varchar(255);NOT NULL;"`
	Description   string    `json:"description" gorm:"type:text;NOT NULL;"`
	CreateAt      time.Time `json:"create_at" gorm:"type:date;NOT NULL;DEFAULT:CURRENT_TIMESTAMP;"`
	UpdateAt      time.Time `json:"update_at" gorm:"type:date;NOT NULL;DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;"`
	MerchantID    uint      `json:"merchant_id" gorm:"type:int(10);NOT NULL;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CategoryID    uint      `json:"category_id" gorm:"type:int(10);NOT NULL;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
