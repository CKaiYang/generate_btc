package main

import (
	"fmt"
	"generate_btc/config"
	"generate_btc/router"
	"generate_btc/util"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.Configure(router.Configure)
	port := config.C.Server.Port
	env := config.C.BTC.ENV
	util.SetNet(env)
	app.Listen(fmt.Sprintf(":%d", port))
}
