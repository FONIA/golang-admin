package lib

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: MD5V
//@description: md5加密
//@param: str []byte
//@return: string
// 大写
func MD5(str []byte) string {
	h := md5.New()
	h.Write(str)
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}