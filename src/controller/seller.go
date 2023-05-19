package controller

import (
	"net/http"
	"rakamin-academy/sdk/message"
	"rakamin-academy/src/model"

	"github.com/gin-gonic/gin"
)

func (c *Controller) SellerRegister(ctx *gin.Context) {

	var registerInput model.RegisterRequest

	if err := ctx.ShouldBindJSON(&registerInput); err != nil {
		message.Error(ctx, http.StatusUnprocessableEntity, "Invalid JSON format", err.Error())
	}

	seller, err := c.ss.Register(ctx.Request.Context(), registerInput)

	if err != nil {
		message.Error(ctx, http.StatusInternalServerError, "Failed to Register", err.Error())
	}

	message.Success(ctx, http.StatusOK, "Success to Register", seller)
}
