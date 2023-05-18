package controller

import (
	"github.com/gin-gonic/gin"
)

func (c *Controller) SellerRegister(ctx *gin.Context) {
	c.ss.Register()
}
