package model

type RegisterRequest struct {
	Name       string `json:"name" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Contact    string `json:"contact" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	BirthDate  string `json:"birth_date" binding:"required"`
	Gender     string `json:"gender" binding:"required"`
	About      string `json:"about" binding:"required"`
	Job        string `json:"job" binding:"required"`
	ProvinceID uint   `json:"province_id" binding:"required"`
	CityID     uint   `json:"city_id" binding:"required"`
	IsAdmin    bool   `json:"is_admin" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}
