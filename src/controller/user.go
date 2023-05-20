package controller

import (
	"net/http"
	"rakamin-academy/sdk/message"
	"rakamin-academy/src/model"

	"github.com/gin-gonic/gin"
)

func (c *Controller) UserRegister(ctx *gin.Context) {

	var registerInput model.RegisterRequest

	if err := ctx.ShouldBindJSON(&registerInput); err != nil {
		message.Error(ctx, http.StatusUnprocessableEntity, "Invalid JSON format", err.Error())
	}

	user, err := c.us.Register(ctx.Request.Context(), registerInput)

	if err != nil {
		message.Error(ctx, http.StatusInternalServerError, "Failed to Register", err.Error())
	}

	message.Success(ctx, http.StatusOK, "Success to Register", user)

}

func (c *Controller) UserLogin(ctx *gin.Context) {

	var loginInput model.LoginRequest

	if err := ctx.ShouldBindJSON(&loginInput); err != nil {
		message.Error(ctx, http.StatusUnprocessableEntity, "Invalid JSON format", err.Error())
	}

	user, err := c.us.Login(ctx.Request.Context(), loginInput)

	if err != nil {
		message.Error(ctx, http.StatusInternalServerError, "Failed to Login", err.Error())
	}

	message.Success(ctx, http.StatusOK, "Success to Login", user)

}
