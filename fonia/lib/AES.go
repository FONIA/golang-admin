package lib

import (
	"bytes"
	"crypto/cipher"
	"crypto/aes"
	"encoding/base64"
)
//CBC
 
//填充字符串（末尾）
func PaddingText1(str []byte, blockSize int) []byte {
	//需要填充的数据长度
	paddingCount := blockSize - len(str)%blockSize
	//填充数据为：paddingCount ,填充的值为：paddingCount
	paddingStr := bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)
	newPaddingStr := append(str, paddingStr...)
	//fmt.Println(newPaddingStr)
	return newPaddingStr
}
 
//去掉字符（末尾）
func UnPaddingText1(str []byte) []byte {
	n := len(str)
	count := int(str[n-1])
	newPaddingText := str[:n-count]
	return newPaddingText
}
//---------------DES加密  解密--------------------
func EncyptogAES(str interface{}, key []byte) interface{} {
	src := []byte(str.(string))
	block,err:=aes.NewCipher(key)
	if err!= nil{
		return nil
	}
	src=PaddingText1(src,block.BlockSize())
	blockMode:=cipher.NewCBCEncrypter(block,key)
	blockMode.CryptBlocks(src,src)
	return base64.StdEncoding.EncodeToString(src)
 
}
func DecrptogAES(src interface{},key []byte) interface{} {
	str,err := base64.StdEncoding.DecodeString(src.(string))
	d_src := []byte(str)
	if err !=nil{
		return nil
	}
	block,err:=aes.NewCipher(key)
	if err!= nil{
		return nil
	}
	blockMode:=cipher.NewCBCDecrypter(block,key)
	blockMode.CryptBlocks(d_src,d_src)
	d_src=UnPaddingText1(d_src)
	return string(d_src)
}

//加密
//明文str
//密钥key:=[]byte("12345678abcdefgh")
//加密字串encodestr := base64.StdEncoding.EncodeToString(EncyptogAES([]byte(str),key))

//解密
//加密base64字串：src
//密钥key:=[]byte("12345678abcdefgh")
//mis,_ := base64.StdEncoding.DecodeString(src)
//DecrptogAES(mis,key)
//解密后明文：mis