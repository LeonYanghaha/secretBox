package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
)

//加密
func EnCrypt(orig, key []byte) string {
	//将秘钥中的每个字节累加,通过sum实现orig的加密工作
	sum := 0
	for i := 0; i < len(key); i++ {
		sum += int(key[0])
	}

	//给明文补码
	var pkcs_code = PKCS5Padding(orig, 8)

	//通过秘钥，对补码后的明文进行加密
	for j := 0; j < len(pkcs_code); j++ {
		pkcs_code[j] += byte(sum)
	}
	return base64.StdEncoding.EncodeToString(pkcs_code)
}

//补码
func PKCS5Padding(orig []byte, size int) []byte {
	//计算明文的长度
	length := len(orig)
	padding := size - length % size
	//向byte类型的数组中重复添加padding
	repeats := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(orig, repeats...)
}

//解密
func DeCrypt(text string, key []byte) string {
	orig, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "密文类型错误"
	}
	sum := 0
	for i := 0; i < len(key); i++ {
		sum += int(key[0])
	}

	//解密
	for j := 0; j < len(orig); j++ {
		orig[j] -= byte(sum)
	}

	//去码
	var pkcs_unCode = PKCS5UnPadding(orig)
	return string(pkcs_unCode)
}

//去码
func PKCS5UnPadding(orig []byte) []byte {
	//获取最后一位补码的数字
	var tail = int(orig[len(orig) - 1])
	return orig[:(len(orig) - tail)]
}

func main() {
	key := []byte("asdfghjkl")
	var enCryptCode = EnCrypt([]byte("zqwertyuioqwertyuisdfghjwertyujdfghjdfghxcvbnm"), key)
	fmt.Println("密文是：", enCryptCode)

	var deCryptCode = DeCrypt(enCryptCode, key)
	fmt.Println("明文是：", deCryptCode)
}

