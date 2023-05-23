package controller

import (
	"net/http"
	"rakamin-academy/sdk/message"
	"rakamin-academy/src/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *Controller) CreateCategory(ctx *gin.Context) {

	userID := ctx.MustGet("user").(uint)

	var newCategory model.CreateCategory

	if err := ctx.ShouldBindJSON(&newCategory); err != nil {
		message.Error(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	category, err := c.cs.Create(ctx.Request.Context(), newCategory, userID)

	if err != nil {
		message.Error(ctx, http.StatusInternalServerError, "Failed to create new category", err.Error())
		return
	}

	message.Success(ctx, http.StatusOK, "Success to create new category", category)

}

func (c *Controller) FindAllCategories(ctx *gin.Context) {

	userID := ctx.MustGet("user").(uint)

	categories, err := c.cs.FindAll(ctx.Request.Context(), userID)

	if err != nil {
		message.Error(ctx, http.StatusInternalServerError, "Failed to fetch all categories", err.Error())
		return
	}

	message.Success(ctx, http.StatusOK, "Success to fetch all categories", categories)

}

func (c *Controller) FindCategoryByID(ctx *gin.Context) {

	userID := ctx.MustGet("user").(uint)

	categoryIDStr := ctx.Param("categoryID")

	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)

	if err != nil {
		message.Error(ctx, http.StatusBadRequest, "Invalid category ID", err.Error())
		return
	}

	category, err := c.cs.FindByID(ctx.Request.Context(), userID, uint(categoryID))

	if err != nil {
		message.Error(ctx, http.StatusInternalServerError, "Failed to fetch category", err.Error())
		return
	}

	message.Success(ctx, http.StatusOK, "Success to fetch category", category)

}

func (c *Controller) UpdateCategory(ctx *gin.Context) {

	userID := ctx.MustGet("user").(uint)

	categoryIDStr := ctx.Param("categoryID")

	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)

	if err != nil {
		message.Error(ctx, http.StatusBadRequest, "Invalid category ID", err.Error())
		return
	}

	var updateCategory model.CreateCategory

	if err := ctx.ShouldBindJSON(&updateCategory); err != nil {
		message.Error(ctx, http.StatusBadRequest, "Invalid JSON format", err.Error())
		return
	}

	category, err := c.cs.Update(ctx.Request.Context(), userID, uint(categoryID), updateCategory)

	if err != nil {
		message.Error(ctx, http.StatusInternalServerError, "Failed to update category", err.Error())
		return
	}

	message.Success(ctx, http.StatusOK, "Success to update category", category)

}
