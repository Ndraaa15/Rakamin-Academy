package product

import "time"

type PhotoProduct struct {
	ID        uint      `json:"id" gorm:"type:int(10) UNSIGNED AUTO_INCREMENT;primaryKey;"`
	ProductID uint      `json:"product_id" gorm:"type:int(10);NOT NULL;"`
	URL       string    `json:"url" gorm:"type:varchar(255);NOT NULL;"`
	UpdateAt  time.Time `json:"update_at" gorm:"type:date;NOT NULL;DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;"`
	CreateAt  time.Time `json:"create_at" gorm:"type:date;NOT NULL;DEFAULT:CURRENT_TIMESTAMP;"`
}
