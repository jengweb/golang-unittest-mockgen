package core

import (
	"fmt"
	"main/controllers"
	"main/gateways"
	"main/services"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func SetupRouter() fasthttp.RequestHandler {
	route := router.New()
	usersGateways := gateways.InitUsersGateways()
	usersServices := services.InitUsersService(usersGateways)
	usersControllers := controllers.InitUsersController(usersServices)

	route.GET("/", func(ctx *fasthttp.RequestCtx) {
		fmt.Println("Hello FastHttp")
	})
	route.GET("/get_users", func(ctx *fasthttp.RequestCtx) {
		usersControllers.GetUsers(ctx)
	})

	return route.Handler
}
