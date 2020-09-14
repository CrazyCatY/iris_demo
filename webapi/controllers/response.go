/*
 * @FilePath: \webapi\controllers\response.go
 * @Descripttion:
 *
 * @Date: 2020-07-23 14:26:42
 * @Author: yuanhao
 *
 */
package controllers

import (
	"net/http"

	"nssa/constants"
	"nssa/util"
	"nssa/webapi/models"
	scode "nssa/webapi/statuscode"

	"github.com/kataras/iris/v12/mvc"
)

const (
// constants.ContentTypeJSON = "application/json"
)

// 获取错误类型返回结构体的字节数组
func NewErrContext(statusCode scode.ErrorCode) []byte {

	if _, ok := scode.StatusCodes[string(statusCode)]; !ok {
		statusCode = "Default"
	}

	resp := models.Response{
		Msg: scode.StatusCodes[string(statusCode)].Msg,
	}

	data, err := util.JsonMarshal(resp)
	if err != nil {
		return nil
	}

	return data
}

func NewOKContext(statusCode scode.ErrorCode, payLoad interface{}) ([]byte, error) {

	if _, ok := scode.StatusCodes[string(statusCode)]; !ok {
		statusCode = "Default"
	}

	resp := models.Response{
		Msg:     scode.StatusCodes[string(statusCode)].Msg,
		Success: true,
		Payload: payLoad,
	}

	return util.JsonMarshal(resp)
}

// NewErrResponse 实例化一个错误返回结构
func NewErrResponse(statusCode scode.ErrorCode) mvc.Response {
	return mvc.Response{
		Code:        scode.StatusCodes[string(statusCode)].Code,
		ContentType: constants.ContentTypeJSON,
		Content:     NewErrContext(statusCode),
	}
}

// NewSuccessResponse 实例化一个OK返回结构
func NewSuccessResponse(payLoad interface{}) mvc.Response {
	context, err := NewOKContext(scode.NILError, payLoad)
	if err != nil {
		return NewJSONMarshalErrResponse()
	}
	return mvc.Response{
		Code:        http.StatusOK,
		ContentType: constants.ContentTypeJSON,
		Content:     context,
	}
}

// NewCreatedResponse 实例化一个创建成功返回结构
func NewCreatedResponse(id string) mvc.Response {

	context, err := NewOKContext(scode.CreateRecordSuccess, id)
	if err != nil {
		return NewJSONMarshalErrResponse()
	}
	return mvc.Response{
		Code:        http.StatusCreated,
		ContentType: constants.ContentTypeJSON,
		Content:     context,
	}
}

// NewUpdateResponse 实例化一个更新成功返回结构
func NewUpdateResponse(id string) mvc.Response {

	context, err := NewOKContext(scode.UpdateRecordSuccess, id)
	if err != nil {
		return NewJSONMarshalErrResponse()
	}

	return mvc.Response{
		Code:        http.StatusOK,
		ContentType: constants.ContentTypeJSON,
		Content:     context,
	}
}

// NewDeleteResponse 实例化一个删除成功返回结构
func NewDeleteResponse(id string) mvc.Response {

	context, err := NewOKContext(scode.DeleteRecordSuccess, id)
	if err != nil {
		return NewJSONMarshalErrResponse()
	}

	return mvc.Response{
		Code:        http.StatusOK,
		ContentType: constants.ContentTypeJSON,
		Content:     context,
	}
}

// NewErrResponse 实例化一个错误返回结构
func NewMsgDataNotExistResponse() mvc.Response {
	return mvc.Response{
		Code:        scode.StatusCodes[string(scode.MsgDataNotExist)].Code,
		ContentType: constants.ContentTypeJSON,
		Content:     NewErrContext(scode.MsgDataNotExist),
	}
}

// NewBadErrResponse 实例化返回一个错误请求返回结构
func NewBadErrResponse(content []byte) mvc.Response {
	return mvc.Response{
		Code:        http.StatusBadRequest,
		ContentType: constants.ContentTypeJSON,
		Content:     content,
	}
}

// NewReadRequestBodyErrorResponse 实例化返回一个读取请求body错误返回结构
func NewReadRequestBodyErrResponse() mvc.Response {
	return mvc.Response{
		Code:        http.StatusBadRequest,
		ContentType: constants.ContentTypeJSON,
		Content:     NewErrContext(scode.ReadRequestBodyError),
	}
}

// NewUpdateErrResponse 实例化返回一个修改数据错误返回结构
func NewUpdateErrResponse() mvc.Response {
	return mvc.Response{
		Code:        http.StatusInternalServerError,
		ContentType: constants.ContentTypeJSON,
		Content:     NewErrContext(scode.UpdateRecordError),
	}
}

// NewDeleteErrResponse 实例化返回一个删除数据错误返回结构
func NewDeleteErrResponse() mvc.Response {
	return mvc.Response{
		Code:        http.StatusInternalServerError,
		ContentType: constants.ContentTypeJSON,
		Content:     NewErrContext(scode.DeleteRecordError),
	}
}

// NewQueryErrResponse 实例化返回一个删除数据错误返回结构
func NewQueryErrResponse() mvc.Response {
	return mvc.Response{
		Code:        http.StatusInternalServerError,
		ContentType: constants.ContentTypeJSON,
		Content:     NewErrContext(scode.QueryRecordError),
	}
}

// NewQueryErrResponse 实例化返回一个删除数据错误返回结构
func NewJSONMarshalErrResponse() mvc.Response {
	return mvc.Response{
		Code:        http.StatusBadRequest,
		ContentType: constants.ContentTypeJSON,
		Content:     NewErrContext(scode.JSONMarshalError),
	}
}

// NewRequestBodyInvalidErrResponse 实例化返回一个读取请求body错误返回结构
func NewRequestBodyInvalidErrResponse() mvc.Response {
	return mvc.Response{
		Code:        scode.StatusCodes[string(scode.RequestBodyInvalid)].Code,
		ContentType: constants.ContentTypeJSON,
		Content:     NewErrContext(scode.RequestBodyInvalid),
	}
}
