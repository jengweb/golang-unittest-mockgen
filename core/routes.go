package core

import (
	"fmt"
	"main/controllers"
	"main/gateways"
	"main/services"

	"github.com/valyala/fasthttp/pprofhandler"

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

	// pprof handler
	route.GET("/debug/pprof/{profile:*}", pprofhandler.PprofHandler)
	// route.Handle("/debug/pprof/", pprof.Index)
	// route.Handle("/debug/pprof/profile", pprof.Profile)
	// route.Handle("/debug/pprof/cmdline", pprof.Cmdline)
	// route.Handle("/debug/pprof/symbol", pprof.Symbol)
	// route.Handle("/debug/pprof/trace", pprof.Trace)

	return route.Handler
}
