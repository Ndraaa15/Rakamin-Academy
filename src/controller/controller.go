package controller

import (
	"fmt"
	"os"
	us "rakamin-academy/src/service/user_service"
	"sync"

	"github.com/gin-gonic/gin"
)

var once = sync.Once{}

type Controller struct {
	http *gin.Engine
	us   us.UserService
}

func NewController(ss us.UserService) *Controller {
	r := Controller{}
	once.Do(func() {
		r.http = gin.Default()
		r.us = ss
		r.RoutesAndMiddleware()
	})
	return &r
}

func (c *Controller) RoutesAndMiddleware() {
	/*
		Todo: Add another routes here
		- Get All Category
		- Get Category by ID
		- Post Category
		- Update Category by ID

		- Login
		- Register

		- Get my Profile
		- Update my Profile

		- Get my Alamat
		- Get Alamat by ID
		- Post Alamat
		- Update Alamat by ID

		- Get Merchant
		- Update merchant Profile
		- Get Merchant by ID
		- Get All Merchant

		- Get All Product
		- Get Product by ID
		- Post Product
		- Update Product by ID
		- Delete Product by ID

		- Get All Transaction
		- Get Transaction by ID

		- Get List Province
		- Get Detail Province
		- Get List City
		- Get Detail City
	*/
}

func (c *Controller) RunServer() {
	c.http.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
