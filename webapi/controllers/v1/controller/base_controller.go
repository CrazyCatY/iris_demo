/*
 * @FilePath: \webapi\controllers\v1\controller\base_controller.go
 * @Descripttion:
 *
 * @Date: 2020-07-23 14:22:18
 * @Author: yuanhao
 *
 */
package controller

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	// jsoniter "github.com/json-iterator/go"
	"github.com/kataras/iris/v12"
)

var (
	// json     = jsoniter.ConfigCompatibleWithStandardLibrary
	validate *validator.Validate
)

func init() {
	validate = validator.New()
}

type baseController struct {
}

//BeginRequest BeginRequest
func (bc *baseController) BeginRequest(ctx iris.Context) {
	//TODO 在HTTP请求开始时做一些事情
}

//EndRequest EndRequest
func (bc *baseController) EndRequest(ctx iris.Context) {
	//TODO 在HTTP请求结束后做一些事情
}

func (bc *baseController) validate(reqBody interface{}) (bool, error) {
	if reqBody != nil && reflect.TypeOf(reqBody).Kind() == reflect.Struct {
		if err := validate.Struct(reqBody); err != nil {
			return false, err
		}
	}

	return true, nil
}
