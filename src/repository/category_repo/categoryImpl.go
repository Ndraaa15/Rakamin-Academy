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

func (c *Category) FindAll(ctx context.Context) ([]e.Category, error) {

	var categories []e.Category

	if err := c.sql.Debug().WithContext(ctx).Find(&categories).Error; err != nil {
		return categories, errors.New("GORM ERROR : " + err.Error())
	}

	return categories, nil

}

func (c *Category) FindByID(ctx context.Context, categoryID uint) (e.Category, error) {

	var category e.Category

	if err := c.sql.Debug().WithContext(ctx).Where("id = ?", categoryID).First(&category).Error; err != nil {
		return category, errors.New("GORM ERROR : " + err.Error())
	}

	return category, nil

}
func (c *Category) Update(ctx context.Context, category e.Category) (e.Category, error) {

	if err := c.sql.Debug().WithContext(ctx).Model(&category).Updates(&category).Error; err != nil {
		return category, errors.New("GORM ERROR : " + err.Error())
	}

	return category, nil

}
