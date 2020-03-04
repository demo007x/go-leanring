package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	clipherStr := h.Sum(nil)

	return hex.EncodeToString(clipherStr)
}

func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// 校验用户密码
func ValidatePasswd(plainpwd, salt, passwd string) bool {
	return Md5Encode(plainpwd+salt) == passwd
}

// 生成密码
func MakePassword(plainpwd, solt string) string {
	return Md5Encode(plainpwd + solt)
}
