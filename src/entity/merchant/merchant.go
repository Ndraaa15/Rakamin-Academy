package merchant

import "time"

type Merchant struct {
	ID       uint `json:"id" gorm:"type:int(10) UNSIGNED AUTO_INCREMENT;primaryKey;"`
	SellerID uint `json:"seller_id" gorm:"type:int(10);NOT NULL;"`
	// Seller    entity.Seller `json:"seller" gorm:"foreignKey:SellerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name     string    `json:"name" gorm:"type:varchar(255);NOT NULL;"`
	URL      string    `json:"url" gorm:"type:varchar(255);NOT NULL;"`
	CreateAt time.Time `json:"create_at" gorm:"type:date;NOT NULL;DEFAULT:CURRENT_TIMESTAMP;"`
	UpdateAt time.Time `json:"update_at" gorm:"type:date;NOT NULL;DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;"`
	// Product []entity.Product `json:"product" gorm:"foreignKey:MerchantID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

}
