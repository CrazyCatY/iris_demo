/*
 * @FilePath: \util\psw_bcrypt.go
 * @Descripttion:
 *
 * @Date: 2020-07-29 14:48:34
 * @Author: yuanhao
 *
 */
package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"

	"github.com/dlclark/regexp2"
	"golang.org/x/crypto/bcrypt"
)

var (
	key = []byte("1234123412ABCDEF1234123412ABCDEF")
)

// 获取经过bcryptb加密后的字符串
func GetBcryptPasswordString(password string) (string, error) {
	hashPsw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashPsw), err
}

// 验证密码是否一致
func CheckPassword(password, hashPsw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPsw), []byte(password)) == nil
}

//AES CSC 加密
func AesEncryptCBC(origData []byte) (encrypted []byte) {
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize)                // 补全码
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式
	encrypted = make([]byte, len(origData))                     // 创建数组
	blockMode.CryptBlocks(encrypted, origData)                  // 加密
	return encrypted
}

//AES CSC 解密
func AesDecryptCBC(encrypted []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)                              // 分组秘钥
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	decrypted = make([]byte, len(encrypted))                    // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)                 // 解密
	decrypted = pkcs5UnPadding(decrypted)                       // 去除补全码

	defer func() {
		_ = recover()
	}()

	return decrypted
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// 检查密码复杂度 必须有数字，大小写，可以有特殊字符，长度【5，255】
func CheckPasswordComplexity(password string) (bool, error) {
	passwordRegex, _ := regexp2.Compile(`^(?![0-9]+$)(?![a-z]+$)(?![A-Z]+$)[0-9A-Za-z!@#\$ %\^&\*\(\)_\-=\+~\[\]\{\}\|;:'"\?/<>,.\\]{8,255}$`, 0)
	return passwordRegex.MatchString(password)
}

// 检查用户名复杂度 用户名位字母加数字、字母，长度【5，255】
func CheckNameComplexity(name string) (bool, error) {
	letterRegex, _ := regexp2.Compile(`^(?=.*[a-zA-Z])[a-zA-Z0-9]{5,255}$`, 0)
	return letterRegex.MatchString(name)
}
