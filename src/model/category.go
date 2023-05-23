package model

type CreateCategory struct {
	Name string `json:"category_name" gorm:"type:varchar(255);NOT NULL;" binding:"required"`
}
