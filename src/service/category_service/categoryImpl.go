package category_service

import (
	"context"
	"errors"
	e "rakamin-academy/src/entity/category"
	"rakamin-academy/src/model"
	cr "rakamin-academy/src/repository/category_repo"
	ur "rakamin-academy/src/repository/user_repo"
)

type Category struct {
	cr cr.CategoryRepository
	ur ur.UserRepository
}

func NewCategoryService(cr cr.CategoryRepository, ur ur.UserRepository) CategoryService {
	return &Category{cr, ur}
}

func (c *Category) Create(ctx context.Context, newCategory model.CreateCategory, userID uint) (e.Category, error) {

	var category e.Category

	user, err := c.ur.FindByID(ctx, userID)

	if err != nil {
		return category, errors.New("SERVICE ERROR : FAILED TO GET USER")
	}

	if !user.IsAdmin {
		return category, errors.New("SERVICE ERROR : USER IS NOT ADMIN")
	}

	category = e.Category{
		Name: newCategory.Name,
	}

	category, err = c.cr.Create(ctx, category)

	if err != nil {
		return category, errors.New("SERVICE ERROR : FAILED TO CREATE NEW CATEGORY")
	}

	return category, nil

}
