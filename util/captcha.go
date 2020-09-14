/*
 * @FilePath: \util\captcha.go
 * @Descripttion:
 *
 * @Date: 2020-07-28 16:55:56
 * @Author: yuanhao
 *
 */
package util

import (
	"image/color"

	"github.com/mojocn/base64Captcha"
)

type CaptchaImage struct {
	ID   string `json:"id"`
	B64s string `json:"b64s"`
}

var (
	driver = base64Captcha.NewDriverString(
		60,
		240,
		0,
		base64Captcha.OptionShowSlimeLine,
		5,
		"1234567890qwertyuioplkjhgfdsazxcvbnm",
		&color.RGBA{R: 254, G: 254, B: 254, A: 254},
		[]string{})
	store   = base64Captcha.DefaultMemStore
	captcha = base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
)

// GenerateCaptcha 生成验证码
func GenerateCaptcha() (id, b64s string, err error) {
	id, b64s, err = captcha.Generate()
	return
}

// VerifyCaptcha 校验验证码
func VerifyCaptcha(id, answer string) bool {
	result := store.Verify(id, answer, true)
	return result
}
