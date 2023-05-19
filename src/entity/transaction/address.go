package transaction

import (
	u "rakamin-academy/src/entity/user"
	"time"
)

type Address struct {
	ID             uint      `json:"id" gorm:"type:int(10) UNSIGNED AUTO_INCREMENT;primaryKey;"`
	SellerID       uint      `json:"seller_id" gorm:"type:char(36);NOT NULL;"`
	Seller         u.User    `json:"seller" gorm:"foreignKey:SellerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AddressTitle   string    `json:"address_title" gorm:"type:varchar(255);NOT NULL;"`
	RecipientName  string    `json:"recipient_name" gorm:"type:varchar(255);NOT NULL;"`
	RecipientPhone string    `json:"recipient_phone" gorm:"type:varchar(255);NOT NULL;"`
	DetailAddress  string    `json:"detail_address" gorm:"type:varchar(255);NOT NULL;"`
	UpdateAt       time.Time `json:"update_at" gorm:"type:date;NOT NULL;DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;"`
	CreateAt       time.Time `json:"create_at" gorm:"type:date;NOT NULL;DEFAULT:CURRENT_TIMESTAMP;"`
}
