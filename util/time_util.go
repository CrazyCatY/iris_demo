/*
 * @FilePath: \util\time_util.go
 * @Descripttion:
 *
 * @Date: 2020-07-30 10:45:18
 * @Author: yuanhao
 *
 */
package util

import "time"

func GetMilliSecondTime() int64 {
	return time.Now().UnixNano() / 1e6
}
