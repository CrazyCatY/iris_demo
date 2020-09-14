/*
 * @FilePath: \webapi\middleware\preflight.go
 * @Descripttion:
 *
 * @Date: 2020-08-03 15:45:02
 * @Author: yuanhao
 *
 */
package middleware

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

//Preflight 应对Ajax预检请求
func Preflight(ctx iris.Context) {
	if ctx.Method() == http.MethodOptions {
		ctx.ResponseWriter().WriteHeader(http.StatusOK)
		return
	}
	ctx.Next()
}
