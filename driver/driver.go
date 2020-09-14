/*
 * @FilePath: \driver\driver.go
 * @Descripttion:
 *
 * @Date: 2020-07-23 16:21:20
 * @Author: yuanhao
 *
 */
package driver

import (
	"sync"

	"nssa/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	once     sync.Once
	instance *gorm.DB
)

// 获取数据库连接（postgres）
func GetDriverDB() *gorm.DB {

	var err error

	once.Do(func() {
		conf := config.GetConfigure()
		instance, err = gorm.Open("postgres", "postgres://"+
			conf.Postgres.UserName+":"+
			conf.Postgres.Password+"@"+
			conf.Postgres.Host+"/"+
			conf.Postgres.DB+
			"?sslmode=disable")

		// 设置连接池
		instance.DB().SetMaxIdleConns(50)
		instance.DB().SetMaxOpenConns(500)
		// 设置日志模式
		instance.LogMode(conf.Postgres.Log)

	})

	if err != nil {
		return nil
	}

	return instance
}
