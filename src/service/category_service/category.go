package category_service

import (
	"context"
	e "rakamin-academy/src/entity/category"
	"rakamin-academy/src/model"
)

type CategoryService interface {
	Create(ctx context.Context, newCategory model.CreateCategory, userID uint) (e.Category, error)
	FindAll(ctx context.Context, userID uint) ([]e.Category, error)
	FindByID(ctx context.Context, userID uint, categoryID uint) (e.Category, error)
	Update(ctx context.Context, userID uint, categoryID uint, updateCategory model.CreateCategory) (e.Category, error)
}
