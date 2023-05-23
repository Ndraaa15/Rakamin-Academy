package model

type CreateCategory struct {
	Name string `json:"category" gorm:"type:varchar(255);NOT NULL;" binding:"required"`
}

type UpdateCategory struct {
	Name string `json:"category" gorm:"type:varchar(255);NOT NULL;" binding:"required"`
}
