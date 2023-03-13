package util

import (
	"generate_btc/models"
	"github.com/kataras/iris/v12"
)

func CommonResultSuccess(ctx iris.Context, data interface{}) {
	ctx.JSON(models.Success(data))
}

func CommonResultSuccess2(ctx iris.Context) {
	ctx.JSON(models.Success(iris.Map{}))
}

func CommonResultFailure(ctx iris.Context, code int, msg string) {
	ctx.JSON(models.Failure(code, msg))
}
