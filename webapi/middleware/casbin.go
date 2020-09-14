/*
 * @FilePath: \webapi\middleware\casbin.go
 * @Descripttion:
 *
 * @Date: 2020-07-29 10:06:11
 * @Author: yuanhao
 *
 */
package middleware

import (
	"net/http"
	"nssa/constants"
	"nssa/webapi/models"
	scode "nssa/webapi/statuscode"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/kataras/iris/v12"
)

var (
	enforcer *casbin.Enforcer
)

func init() {
	initCasbin()
}

func Casbin(ctx iris.Context) {

	// 检查权限
	ok, err := checkCasbin(ctx.Values().GetString(constants.DefaultContextUserRole), ctx.Path(), ctx.Method())
	if err != nil {
		return
	}
	// 权限不通过，返回错误信息
	if !ok {
		ctx.StatusCode(http.StatusForbidden)
		ctx.ContentType(constants.ContentTypeJSON)
		resp := models.Response{
			Msg: scode.StatusCodes[string(scode.RequestForbidden)].Msg,
		}

		if _, err := ctx.JSON(resp); err != nil {
			return
		}
		return
	}

	ctx.Next()
}

// 初始化casbin
func initCasbin() {

	modelText := `
	[request_definition]
	r = sub, obj, act
	
	[policy_definition]
	p = sub, obj,act

	[role_definition]
	g = _, _
	
	[policy_effect]
	e = some(where (p.eft == allow))
	
	[matchers]
	m = g(r.sub, p.sub) && keyMatch(r.obj, p.obj) && (r.act == p.act || p.act == "*")
	`
	m, _ := model.NewModelFromString(modelText)
	enforcer, _ = casbin.NewEnforcer(m)

	setCasbin()
}

// 设置权限
func setCasbin() {
	_, _ = enforcer.AddPermissionForUser("admin", "/api/v1/system/lists", "GET")
}

// 检查权限
func checkCasbin(role, path, method string) (bool, error) {
	return enforcer.Enforce(role, path, method)
}
