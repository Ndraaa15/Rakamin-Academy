package category_repo

import (
	"context"
	"errors"
	"rakamin-academy/database"
	e "rakamin-academy/src/entity/category"
)

type Category struct {
	sql      *database.SQL
	supabase *database.Supabase
}

func NewCategoryRepository(sql *database.SQL, supabase *database.Supabase) CategoryRepository {
	return &Category{sql, supabase}
}

func (c *Category) Create(ctx context.Context, category e.Category) (e.Category, error) {

	if err := c.sql.Debug().WithContext(ctx).Create(&category).Error; err != nil {
		return category, errors.New("GORM ERROR : " + err.Error())
	}

	return category, nil

}
