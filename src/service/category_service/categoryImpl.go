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

func (c *Category) FindAll(ctx context.Context, userID uint) ([]e.Category, error) {

	user, err := c.ur.FindByID(ctx, userID)

	if err != nil {
		return []e.Category{}, errors.New("SERVICE ERROR : FAILED TO GET USER")
	}

	if !user.IsAdmin {
		return []e.Category{}, errors.New("SERVICE ERROR : USER IS NOT ADMIN")
	}

	var categories []e.Category

	categories, err = c.cr.FindAll(ctx)

	if err != nil {
		return categories, errors.New("SERVICE ERROR : FAILED TO GET ALL CATEGORIES")
	}

	return categories, nil

}

func (c *Category) FindByID(ctx context.Context, userID uint, categoryID uint) (e.Category, error) {

	user, err := c.ur.FindByID(ctx, userID)

	if err != nil {
		return e.Category{}, errors.New("SERVICE ERROR : FAILED TO GET USER")
	}

	if !user.IsAdmin {
		return e.Category{}, errors.New("SERVICE ERROR : USER IS NOT ADMIN")
	}

	var category e.Category

	category, err = c.cr.FindByID(ctx, categoryID)

	if err != nil {
		return category, errors.New("SERVICE ERROR : FAILED TO GET CATEGORY")
	}

	return category, nil

}

func (c *Category) Update(ctx context.Context, userID uint, categoryID uint, updateCategory model.CreateCategory) (e.Category, error) {

	user, err := c.ur.FindByID(ctx, userID)

	if err != nil {
		return e.Category{}, errors.New("SERVICE ERROR : FAILED TO GET USER")
	}

	if !user.IsAdmin {
		return e.Category{}, errors.New("SERVICE ERROR : USER IS NOT ADMIN")
	}

	category, err := c.cr.FindByID(ctx, categoryID)

	if err != nil {
		return category, errors.New("SERVICE ERROR : FAILED TO GET CATEGORY")
	}

	category.Name = updateCategory.Name

	category, err = c.cr.Update(ctx, category)

	if err != nil {
		return category, errors.New("SERVICE ERROR : FAILED TO UPDATE CATEGORY")
	}

	return category, nil

}
