/*
 * @FilePath: \webapi\middleware\auth.go
 * @Descripttion:
 *
 * @Date: 2020-07-28 11:27:33
 * @Author: yuanhao
 *
 */
package middleware

import (
	"github.com/kataras/iris/v12"
)

var (
	whiteSkipUrls []string
)

func Auth(ctx iris.Context) {

	// skipFun := GetSkipHandle()
	// if skipFun(ctx.Method() + "-" + ctx.Path()) {
	// 	ctx.Next()
	// 	return
	// }

	// // 获取token
	// token := ctx.GetHeader(constants.DefaultTokenHeader)

	// // 解析token
	// claims, ok := util.JWTTokenIsValid(token)
	// if !ok {
	// 	ctx.StatusCode(http.StatusUnauthorized)
	// 	ctx.ContentType(constants.ContentTypeJSON)
	// 	resp := models.Response{
	// 		Msg: scode.StatusCodes[string(scode.AccessTokenInvalid)].Msg,
	// 	}

	// 	if _, err := ctx.JSON(resp); err != nil {
	// 		return
	// 	}

	// 	return
	// }

	// // 获取用户信息 redis
	// data, _ := redis.GetRedisValue(claims.UUID)
	// user := models.RedisUser{}
	// _ = util.JsonUnmarshal([]byte(data), &user)

	// // 保存用户基本信息
	// ctx.Values().Set(constants.DefaultContextUserId, user.ID)
	// ctx.Values().Set(constants.DefaultContextUserRole, user.Role)

	ctx.Next()
}

func GetSkipHandle() func(string) bool {

	// whiteSkipUrls = append(whiteSkipUrls, iris.MethodGet+"-"+"/api/v1/task/ws/window_logs")
	whiteSkipUrls = append(whiteSkipUrls, iris.MethodGet+"-"+"/api/v1/task/ws")

	return func(urlPath string) bool {

		pathLen := len(urlPath)
		for _, prefix := range whiteSkipUrls {
			if prefixLen := len(prefix); pathLen >= prefixLen && urlPath[:prefixLen] == prefix {
				return true
			}
		}
		return false
	}
}

func WSAuth(ctx iris.Context) {

	// // 获取token
	// token := ctx.GetHeader(constants.DefaultTokenHeader)

	// // 解析token
	// claims, ok := util.JWTTokenIsValid(token)
	// if !ok {
	// 	ctx.StatusCode(http.StatusUnauthorized)
	// 	ctx.ContentType(constants.ContentTypeJSON)
	// 	resp := models.Response{
	// 		Msg: scode.StatusCodes[string(scode.AccessTokenInvalid)].Msg,
	// 	}

	// 	if _, err := ctx.JSON(resp); err != nil {
	// 		return
	// 	}

	// 	return
	// }

	// // 获取用户信息 redis
	// data, _ := redis.GetRedisValue(claims.UUID)
	// user := models.RedisUser{}
	// _ = util.JsonUnmarshal([]byte(data), &user)

	// // 保存用户基本信息
	// ctx.Values().Set(constants.DefaultContextUserId, user.ID)
	// ctx.Values().Set(constants.DefaultContextUserRole, user.Role)
	ctx.Next()
}
