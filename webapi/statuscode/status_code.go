/*
 * @FilePath: \webapi\statuscode\status_code.go
 * @Descripttion:
 *
 * @Date: 2020-07-23 14:44:35
 * @Author: yuanhao
 *
 */
package statuscode

import "net/http"

type StatusResponse struct {
	Code int
	Msg  string
}

var StatusCodes = map[string]StatusResponse{
	"ReadRequestBodyError": StatusResponse{http.StatusBadRequest, "HTTP Request Body Read Error"},
	"JSONMarshalError":     StatusResponse{http.StatusBadRequest, "JSON Marshal Error"},
	"JSONUnmarshalError":   StatusResponse{http.StatusBadRequest, "JSON Unmarshal Error"},
	"AccessDBError":        StatusResponse{http.StatusInternalServerError, "Access Database Error"},
	"AccessTokenInvalid":   StatusResponse{http.StatusUnauthorized, "Access Token Invalid"},
	"RequestBodyInvalid":   StatusResponse{http.StatusBadRequest, "Request Body Invalid"},
	"RequestForbidden":     StatusResponse{http.StatusForbidden, "Request Forbidden"},
	"LoginRecordError":     StatusResponse{http.StatusUnauthorized, "Login Error"},

	// common db error
	"CreateRecordError":    StatusResponse{http.StatusInternalServerError, "Create Record Error"},
	"DeleteRecordError":    StatusResponse{http.StatusInternalServerError, "Delete Record Error"},
	"UpdateRecordError":    StatusResponse{http.StatusInternalServerError, "Update Record Error"},
	"QueryRecordError":     StatusResponse{http.StatusInternalServerError, "Query Record Error"},
	"MsgDataNotExist":      StatusResponse{http.StatusOK, "Msg Data Not Exist"},
	"BatchInsertDataError": StatusResponse{http.StatusInternalServerError, "Batch Insert Data Error"},

	// common error
	"Default":               StatusResponse{http.StatusInternalServerError, "Unknown Error"},
	"OK":                    StatusResponse{http.StatusOK, "OK"},
	"InformationIncomplete": StatusResponse{http.StatusNotFound, "Basic Information Incomplete"},
	"DataFormatError":       StatusResponse{http.StatusNotFound, "Data Format Error"},
	"ReadParametersError":   StatusResponse{http.StatusNotFound, "Read Parameters Error"},

	// success code
	"CreateRecordSuccess": StatusResponse{http.StatusCreated, "Create Record Success"},
	"DeleteRecordSuccess": StatusResponse{http.StatusOK, "Delete Record Success"},
	"UpdateRecordSuccess": StatusResponse{http.StatusOK, "Update Record Success"},
	"QueryRecordSuccess":  StatusResponse{http.StatusOK, "Query Record Success"},
	"IPHasExist":          StatusResponse{http.StatusBadRequest, "IP Has Exist"},

	// token and captcha
	"CreateTokenError":      StatusResponse{http.StatusInternalServerError, "Create Token Error"},
	"CreateCaptchaError":    StatusResponse{http.StatusInternalServerError, "Create Captcha Error"},
	"CaptchaWrongError":     StatusResponse{http.StatusBadRequest, "Captcha Wrong Error"},
	"WrongPassowrdError":    StatusResponse{http.StatusUnauthorized, "Wrong Password Error"},
	"DncryptPasswordError":  StatusResponse{http.StatusInternalServerError, "Dncrypt Password Fail"},
	"AesDecryptCBCPSWError": StatusResponse{http.StatusBadRequest, "Aesdecrypt CBC Password Error"},
	"PasswordSimpleError":   StatusResponse{http.StatusBadRequest, "Password too Simple Error"},
	"UserNameSimpleError":   StatusResponse{http.StatusBadRequest, "UserName too Simple Error"},

	// redis error
	"RedisGetOrSetError": StatusResponse{http.StatusInternalServerError, "Redis Get or Set Data Error"},

	// user error
	"UserNameExistError": StatusResponse{http.StatusBadRequest, "User Name Exists"},
	"UserLoginError":     StatusResponse{http.StatusBadRequest, "User Not Exist Or User Status Disabled"},

	// file error
	"UploadFileReadError": StatusResponse{http.StatusBadRequest, "upload File Read Fail"},
	"UploadFileOpenError": StatusResponse{http.StatusBadRequest, "upload File open Fail"},
}

type ErrorCode string

const (
	UnknownError ErrorCode = "Default"
	NILError     ErrorCode = "OK"

	// ReadRequestBodyError HTTP 请求体读取错误
	ReadRequestBodyError ErrorCode = "ReadRequestBodyError"
	// JSONMarshalError JSON序列化错误
	JSONMarshalError ErrorCode = "JSONMarshalError"
	// JSONUnmarshalError JSON反序列化错误
	JSONUnmarshalError ErrorCode = "JSONUnmarshalError"
	// AccessDBError 访问数据库错误
	AccessDBError ErrorCode = "AccessDBError"
	// UnExistEntityID 不存在这个实体
	UnExistEntityID ErrorCode = "UnExistEntityID"
	// CreateRecordError 新建数据出错
	CreateRecordError ErrorCode = "CreateRecordError"
	// DeleteRecordError 删除数据错误
	DeleteRecordError ErrorCode = "DeleteRecordError"
	// UpdateRecordError 修改数据出错
	UpdateRecordError ErrorCode = "UpdateRecordError"
	// QueryRecordError 获取数据错误
	QueryRecordError ErrorCode = "QueryRecordError"
	// AccessTokenValid 访问令牌无效
	AccessTokenInvalid ErrorCode = "AccessTokenInvalid"
	// RequestForbidden 禁止访问
	RequestForbidden ErrorCode = "RequestForbidden"
	// RequestBodyValid Request Body无效
	RequestBodyInvalid ErrorCode = "RequestBodyInvalid"
	// 基础信息未配置完全
	InformationIncomplete ErrorCode = "InformationIncomplete"
	// 登录失败
	LoginRecordError ErrorCode = "LoginRecordError"
	// 数据转化错误
	DataFormatError ErrorCode = "DataFormatError"
	// Parameters读取错误
	ReadParametersError ErrorCode = "ReadParametersError"

	// 请求成功信息
	// CreateRecordSuccess 新建数据出错
	CreateRecordSuccess ErrorCode = "CreateRecordSuccess"
	// DeleteRecordError 删除数据错误
	DeleteRecordSuccess ErrorCode = "DeleteRecordSuccess"
	// UpdateRecordError 修改数据出错
	UpdateRecordSuccess ErrorCode = "UpdateRecordSuccess"
	// QueryRecordError 获取数据错误
	QueryRecordSuccess ErrorCode = "QueryRecordSuccess"
	// MsgDataNotExist
	MsgDataNotExist ErrorCode = "MsgDataNotExist"
	// MsgDataHasExist
	IPHasExist ErrorCode = "IPHasExist"
	// RedisGetOrSetError
	RedisGetOrSetError ErrorCode = "RedisGetOrSetError"

	// CreateTokenError 生成token失败
	CreateTokenError ErrorCode = "CreateTokenError"
	// CreateCaptchaError 生成验证码失败
	CreateCaptchaError ErrorCode = "CreateCaptchaError"
	// CaptchaWrongError 填写的验证码错误
	CaptchaWrongError ErrorCode = "CaptchaWrongError"
	// WrongPassowrdError 登录密码错误
	WrongPassowrdError ErrorCode = "WrongPassowrdError"
	// DncryptPasswordError 加密密码失败
	DncryptPasswordError ErrorCode = "DncryptPasswordError"
	// AesDecryptCBCPSWError AES解密失败
	AesDecryptCBCPSWError ErrorCode = "AesDecryptCBCPSWError"
	// PasswordSimpleError 密码太简单错误
	PasswordSimpleError ErrorCode = "PasswordSimpleError"
	// UserNameSimpleError 用户名太简单错误
	UserNameSimpleError ErrorCode = "UserNameSimpleError"

	// UserNameExistError 用户名已存在
	UserNameExistError ErrorCode = "UserNameExistError"
	// UserLoginError 用户登录错误
	UserLoginError ErrorCode = "UserLoginError"

	// UploadFileReadError 获取上传文件失败
	UploadFileReadError ErrorCode = "UploadFileReadError"
	// UploadFileOpenError 打开上传文件失败
	UploadFileOpenError ErrorCode = "UploadFileOpenError"
	// BatchInsertDataError 批量插入数据库数据失败
	BatchInsertDataError ErrorCode = "BatchInsertDataError"
)
