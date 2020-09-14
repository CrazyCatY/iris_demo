/*
 * @FilePath: \webapi\iris.go
 * @Descripttion:
 *
 * @Date: 2020-07-23 11:44:13
 * @Author: yuanhao
 *
 */

package webapi

import (
	"fmt"

	"nssa/log"

	v1 "nssa/webapi/controllers/v1"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"

	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"

	_ "nssa/docs" // docs is generated by Swag CLI, you have to import it.
)

// 运行app服务
func RunApp(port int) {

	err := newApp().Run(iris.Addr(fmt.Sprintf(":%d", port)))
	if err != nil {
		log.Printf("Web Server %s", err.Error())
	}
}

// 创建新的服务app
func newApp() *iris.Application {
	app := iris.New()

	app.Use(recover.New())
	app.Use(logger.New())
	//允许跨域
	app.Use(cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:     []string{"*"},
		AllowCredentials:   true,
		OptionsPassthrough: true,
	}))

	// 注册路由
	v1.RegisterRouter(app)

	config := &swagger.Config{
		URL: "http://localhost:8089/swagger/doc.json", //The url pointing to API definition
	}
	// use swagger middleware to
	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(config, swaggerFiles.Handler))

	return app
}