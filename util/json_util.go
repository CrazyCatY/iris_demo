/*
 * @FilePath: \util\json_util.go
 * @Descripttion:
 *
 * @Date: 2020-07-29 16:37:25
 * @Author: yuanhao
 *
 */
package util

import (
	jsoniter "github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func JsonMarshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func JsonUnmarshal(bytes []byte, i interface{}) error {
	return json.Unmarshal(bytes, i)
}
