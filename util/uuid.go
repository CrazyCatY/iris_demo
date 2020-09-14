/*
 * @FilePath: \util\uuid.go
 * @Descripttion:
 *
 * @Date: 2020-07-27 11:49:24
 * @Author: yuanhao
 *
 */
package util

import (
	uuid "github.com/satori/go.uuid"
)

func GetUUID() string {
	return uuid.NewV4().String()
}
