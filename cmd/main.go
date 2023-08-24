package main

import (
	"fmt"
	"main/core"

	"github.com/valyala/fasthttp"
)

func main() {
	fmt.Println("<<== Start Server ==>>")
	// db := core.SetupDatabase()
	routes := core.SetupRouter()
	fasthttp.ListenAndServe(":8000", routes)
}
