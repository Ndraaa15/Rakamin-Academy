package category_repo

import (
	"context"
	e "rakamin-academy/src/entity/category"
)

type CategoryRepository interface {
	Create(ctx context.Context, category e.Category) (e.Category, error)
	FindAll(ctx context.Context) ([]e.Category, error)
	FindByID(ctx context.Context, categoryID uint) (e.Category, error)
	Update(ctx context.Context, category e.Category) (e.Category, error)
}
