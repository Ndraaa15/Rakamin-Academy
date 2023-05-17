package seller

import (
	"time"
)

type Seller struct {
	ID         uint      `json:"id" gorm:"type:int(10) UNSIGNED AUTO_INCREMENT;primaryKey;"`
	Name       string    `json:"name" gorm:"type:varchar(255);NOT NULL;"`
	Password   string    `json:"password" gorm:"type:varchar(255);NOT NULL;"`
	Contact    string    `json:"contact" gorm:"type:varchar(255);NOT NULL;UNIQUE;"`
	Email      string    `json:"email" gorm:"type:varchar(255);NOT NULL;UNIQUE;"`
	BirthDate  time.Time `json:"birth_date" gorm:"type:date;NOT NULL;"`
	Gender     string    `json:"gender" gorm:"type:varchar(255);NOT NULL;"`
	About      string    `json:"about" gorm:"type:text;NOT NULL;"`
	Job        string    `json:"job" gorm:"type:varchar(255);NOT NULL;"`
	ProvinceID uint      `json:"province_id" gorm:"type:varchar(255);NOT NULL;"`
	CityID     uint      `json:"city_id" gorm:"type:varchar(255);NOT NULL;"`
	IsAdmin    bool      `json:"is_admin" gorm:"type:tinyint(1);NOT NULL;DEFAULT:0;"`
	UpdateAt   time.Time `json:"update_at" gorm:"type:date;NOT NULL;DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;"`
	CreateAt   time.Time `json:"create_at" gorm:"type:date;NOT NULL;DEFAULT:CURRENT_TIMESTAMP;"`
}
