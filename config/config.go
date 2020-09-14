/*
 * @FilePath: \config\config.go
 * @Descripttion:
 *
 * @Date: 2020-07-23 16:01:58
 * @Author: yuanhao
 *
 */
package config

import (
	"io/ioutil"
	"sync"

	"github.com/BurntSushi/toml"
)

var (
	once sync.Once
	conf Configure
)

// 配置文件结构体
type Configure struct {
	Server struct {
		Port int `toml:"port"`
	} `toml:"server"`

	Postgres struct {
		Host         string `toml:"host"`
		UserName     string `toml:"username"`
		Password     string `toml:"password"`
		DB           string `toml:"db"`
		MaxOpenConns int    `toml:"max_open_conns"`
		MaxIdleConns int    `toml:"max_idle_conns"`
		Log          bool   `toml:"log"`
	} `toml:"postgres"`

	Redis struct {
		Host     string `toml:"host"`
		Password string `toml:"password"`
	} `toml:"redis"`

	Kafka struct {
		Broker string `toml:"broker"`
	} `toml:"kafka"`
}

// 获取配置文件
func GetConfigure() *Configure {

	once.Do(func() {
		tomlData, err := ioutil.ReadFile("conf/config.toml")
		if err != nil {
			panic(err)
		}

		if _, err := toml.Decode(string(tomlData), &conf); err != nil {
			panic(err)
		}
	})
	return &conf
}
