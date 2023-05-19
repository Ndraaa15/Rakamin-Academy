package merchant

import (
	p "rakamin-academy/src/entity/product"
	u "rakamin-academy/src/entity/user"
	"time"
)

type Merchant struct {
	ID       uint        `json:"id" gorm:"type:int(10) UNSIGNED AUTO_INCREMENT;primaryKey;"`
	SellerID uint        `json:"seller_id" gorm:"type:int(10);NOT NULL;"`
	Seller   u.User      `json:"seller" gorm:"foreignKey:SellerID;"`
	Name     string      `json:"name" gorm:"type:varchar(255);NOT NULL;"`
	URL      string      `json:"url" gorm:"type:varchar(255);NOT NULL;"`
	CreateAt time.Time   `json:"create_at" gorm:"type:date;NOT NULL;DEFAULT:CURRENT_TIMESTAMP;"`
	UpdateAt time.Time   `json:"update_at" gorm:"type:date;NOT NULL;DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;"`
	Product  []p.Product `json:"product" gorm:"foreignKey:MerchantID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
