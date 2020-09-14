/*
 * @FilePath: \webapi\controllers\v1\controller\login_controller.go
 * @Descripttion:
 *
 * @Date: 2020-07-28 13:39:59
 * @Author: yuanhao
 *
 */
package controller

import (
	"nssa/webapi/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// 登录模块控制器
type LoginController struct {
	baseController
}

func NewLoginController() *LoginController {
	return &LoginController{}
}

// @Tags 用户模块
// @Summary 用户登录
// @Description 用户名、密码和验证码登录
// @Accept  json
// @Produce  json
// @Success 200 {string} string	{"msg": "OK","success": true"payload": null}
// @Failure 400 {string} string	{"msg": "OK","success": true"payload": null}
// @Failure 401 {string} string	{"msg": "OK","success": true"payload": null}
// @Failure 403 {string} string	{"msg": "OK","success": true"payload": null}
// @Failure 404 {string} string	{"msg": "OK","success": true"payload": null}
// @Failure 500 {string} string	{"msg": "OK","success": true"payload": null}
// @Router /login [post]
func (lc *LoginController) UserLogin(ctx iris.Context) mvc.Result {

	return controllers.NewSuccessResponse("ok")
}

// 注册登录模块路由
func (lc *LoginController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/login", "UserLogin")
}
