package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// AES加密
// PKCS7 填充模式
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个,然后合并成新的字节切片返回
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7填充的反向操作，删除填充的字符串
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	paddLen := int(origData[length-1])

	end := length - paddLen

	text := origData[:length]
	if end > 0 {
		text = origData[:end]
	}

	return text
}

// aes cbc加密操作
func AesEncrypt(origData []byte, key []byte) ([]byte, error) {

	key, _ = base64.StdEncoding.DecodeString(string(key))

	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块大小
	blockSize := block.BlockSize()
	//对数据进行填充，让数据长度满足需求
	origData = PKCS7Padding(origData, blockSize)

	iv := make([]byte, blockSize)

	//采用AES加密方法中的CBC加密模式
	blocMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	//执行加密
	blocMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// aes cbc解密操作
func AesDecrypt(cypted []byte, key []byte, appointIv ...string) ([]byte, error) {

	key, _ = base64.StdEncoding.DecodeString(string(key))

	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	iv := make([]byte, blockSize)

	if len(appointIv) > 0 {
		iv, _ = base64.StdEncoding.DecodeString(appointIv[0])
	}

	//采用AES加密方法中的CBC加密模式 创建加密客户端实例
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cypted))
	//这个函数还可以用来解密
	blockMode.CryptBlocks(origData, cypted)
	//去除填充字符串
	origData = PKCS7UnPadding(origData)
	return origData, err
}
