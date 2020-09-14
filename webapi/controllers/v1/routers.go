/*
 * @FilePath: \webapi\controllers\v1\routers.go
 * @Descripttion:
 *
 * @Date: 2020-07-23 13:23:54
 * @Author: yuanhao
 *
 */

package controllers

import (
	api "nssa/webapi/controllers/v1/controller"
	"nssa/webapi/middleware"

	iris "github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func RegisterRouter(app *iris.Application) {

	NoAuthParty := app.Party("/api/v1")
	NoAuthParty.Use(middleware.Preflight)
	NoAuthParty.Options("/*", nil)
	mvc.New(NoAuthParty).Handle(api.NewLoginController())
}
