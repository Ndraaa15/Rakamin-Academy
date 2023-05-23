package category_service

import (
	"context"
	e "rakamin-academy/src/entity/category"
	"rakamin-academy/src/model"
)

type CategoryService interface {
	Create(ctx context.Context, newCategory model.CreateCategory, userID uint) (e.Category, error)
}
