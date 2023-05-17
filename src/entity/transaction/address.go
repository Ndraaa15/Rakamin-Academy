package transaction

import (
	"github.com/google/uuid"
)

type Adress struct {
	ID       uuid.UUID `json:"id" gorm:"type:char(36);primaryKey;"`
	SellerID uuid.UUID `json:"seller_id" gorm:"type:char(36);NOT NULL;"`
}
