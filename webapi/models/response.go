/*
 * @FilePath: \webapi\models\response.go
 * @Descripttion:
 *
 * @Date: 2020-07-27 16:09:42
 * @Author: yuanhao
 *
 */
package models

// 返回请求结构体
type Response struct {
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
	Payload interface{} `json:"payload"`
}
