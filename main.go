/*
 * @FilePath: \main.go
 * @Descripttion:
 *
 * @Date: 2020-07-23 11:40:10
 * @Author: yuanhao
 *
 */

package main

import (
	"nssa/config"
	"nssa/driver"
	"nssa/webapi"
)

func init() {
	// 初始化数据库
	db := driver.GetDriverDB()
	if db == nil {
		panic("fail to connect database")
	}

	// 初始化redis
	rdb := driver.GetRedisDB()
	if rdb == nil {
		panic("fail to connect redis")
	}
}

// @title iris_demo
// @version 2.1
// @description iris_demo
// @host 127.0.0.1:8089
// @SecurityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api/v1
func main() {

	conf := config.GetConfigure()
	webapi.RunApp(conf.Server.Port)
}
