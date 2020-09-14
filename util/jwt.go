/*
 * @FilePath: \util\jwt.go
 * @Descripttion:
 *
 * @Date: 2020-07-28 11:46:50
 * @Author: yuanhao
 *
 */
package util

import (
	"crypto/sha256"
	"errors"
	"nssa/log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 密文
const hmacSecret = "sign secret"

// 用户信息结构体
type CustomClaims struct {
	UUID string `json:"uuid"`
	Exp  int64  `json:"exp"`
}

// 获取token
func GetJWTToken(uuid string) (string, error) {
	// 初始化
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaims{
		UUID: uuid,
		Exp:  time.Now().Add(time.Duration(24) * time.Hour).Unix(),
	})

	// 加密
	tokenString, err := token.SignedString(calcHMAC([]byte(hmacSecret)))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析token
func JWTTokenIsValid(tokenString string) (*CustomClaims, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return calcHMAC([]byte(hmacSecret)), nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, true
	} else {
		return nil, false
	}
}

// 检验token是否过期
func (c *CustomClaims) Valid() error {
	if c.Exp < time.Now().Unix() {
		return errors.New("token is expired")
	}
	return nil
}

// 加密密文
func calcHMAC(value []byte) []byte {
	h := sha256.New()
	_, err := h.Write(value)
	if err != nil {
		log.Debug("calculate hmac error: ", err)
	}
	return h.Sum(nil)
}
