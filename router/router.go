package router

import (
	"generate_btc/api"
	"generate_btc/service"
	"generate_btc/util"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
)

func Configure(app *iris.Application) {
	//跨域设置
	crs := CORS("")

	app.Use(crs)

	app.AllowMethods(iris.MethodOptions)

	app.OnErrorCode(iris.StatusNotFound, func(ctx *context.Context) {
		util.CommonResultFailure(ctx, iris.StatusNotFound, iris.ErrNotFound.Error())
	})

	app.OnErrorCode(iris.StatusInternalServerError, func(ctx *context.Context) {
		util.CommonResultFailure(ctx, iris.StatusInternalServerError, "the server has error,please contact us")
	})

	// 应用 context path
	v1 := app.Party("/")
	wallet := mvc.New(v1.Party("/wallet"))
	//依赖注入
	wallet.Register(new(service.WalletService))
	wallet.Handle(new(api.WalletApi))

}

func CORS(allowedOrigin string) iris.Handler { // or "github.com/iris-contrib/middleware/cors"
	if allowedOrigin == "" {
		allowedOrigin = "*"
	}

	return func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", allowedOrigin)
		ctx.Header("Access-Control-Allow-Credentials", "true")
		// July 2021 Mozzila updated the following document: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy
		ctx.Header("Referrer-Policy", "no-referrer-when-downgrade")
		ctx.Header("Access-Control-Expose-Headers", "*, Authorization, X-Authorization")
		if ctx.Method() == iris.MethodOptions {
			ctx.Header("Access-Control-Allow-Methods", "*")
			ctx.Header("Access-Control-Allow-Headers", "*")
			ctx.Header("Access-Control-Max-Age", "86400")
			ctx.StatusCode(iris.StatusNoContent)
			return
		}

		ctx.Next()
	}
}
