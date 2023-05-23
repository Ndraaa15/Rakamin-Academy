package category_repo

import (
	"context"
	e "rakamin-academy/src/entity/category"
)

type CategoryRepository interface {
	Create(ctx context.Context, category e.Category) (e.Category, error)
}
