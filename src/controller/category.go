package controller

import (
	"net/http"
	"rakamin-academy/sdk/message"
	"rakamin-academy/src/model"

	"github.com/gin-gonic/gin"
)

func (c *Controller) CreateCategory(ctx *gin.Context) {

	userID := ctx.MustGet("user").(uint)

	var newCategory model.CreateCategory

	if err := ctx.ShouldBindJSON(&newCategory); err != nil {
		message.Error(ctx, 400, "Invalid JSON format", err.Error())
		return
	}

	category, err := c.cs.Create(ctx.Request.Context(), newCategory, userID)

	if err != nil {
		message.Error(ctx, http.StatusInternalServerError, "Failed to create new category", err.Error())
		return
	}

	message.Success(ctx, http.StatusOK, "Success to create new category", category)

}
