package controller

import (
	"fmt"
	"os"
	ss "rakamin-academy/src/service/seller_service"
	"sync"

	"github.com/gin-gonic/gin"
)

var once = sync.Once{}

type Controller struct {
	http *gin.Engine
	ss   ss.SellerService
}

func NewController(ss ss.SellerService) *Controller {
	r := Controller{}
	once.Do(func() {
		r.http = gin.Default()
		r.ss = ss
		r.RoutesAndMiddleware()

	})
	return &r
}

func (c *Controller) RoutesAndMiddleware() {
	c.http.POST("/seller/register", c.SellerRegister)
}

func (c *Controller) RunServer() {
	c.http.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
