package controllers

import (
	"encoding/json"
	"main/services"

	"github.com/valyala/fasthttp"
)

type UsersController interface {
	GetUsers(ctx *fasthttp.RequestCtx)
}

type usersController struct {
	usersService services.UsersService
}

func InitUsersController(usersService services.UsersService) UsersController {
	return &usersController{
		usersService,
	}
}

func (c *usersController) GetUsers(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("content-type", "application/json")

	resp, err := c.usersService.GetUsers()
	if err != nil {
		ctx.Response.SetStatusCode(500)
		ctx.Error(err.Error(), 500)
		return
	}
	users, _ := json.Marshal(resp)
	ctx.Write(users)
}
