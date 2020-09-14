/*
 * @FilePath: \constants\constant.go
 * @Descripttion:
 *
 * @Date: 2020-07-28 15:33:08
 * @Author: yuanhao
 *
 */
package constants

import (
	"errors"
	"time"
)

const (
	// header token
	DefaultTokenHeader = "Authorization"

	// ContentTypeJSON
	ContentTypeJSON = "application/json"

	// redis用户信息 过期时间
	RedisUserExp = 24 * time.Hour

	//context获取用户id的key值
	DefaultContextUserId = "X-USER-ID"
	//context获取用户role的key值
	DefaultContextUserRole = "X-USER-ROLE"
)

var (
	ErrorUserExist = errors.New("user exists")
	IPHasExist     = errors.New("ip exists")
)
