package toolBox

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"secretBox/models"
)
// 更新secret
func updateSecret(line3 string)(isOk bool){

	fileContent,_ := fileReadByLine(getFilePath())
	if len(fileContent)!=4 {
		return false
	}

	line3 = line3+"\n"
	content:=[]byte(fileContent[0] + fileContent[1] + line3+ fileContent[3])
	err:=writeFile(content,0777)
	if err!=nil {
		fmt.Println(err)
		return false
	}else{
		return true
	}
}
// 保存secret
func saveSecret(secret  models.Secret)(isOk bool){

	secretStr , error := getSecrecList()
	if error != "" {
		return  false
	}
	if secretStr=="\n" {
		secretStr = ""
	}
	var st1 []models.Secret
	if secretStr!="" {
		secretBuf := []byte(secretStr)
		var str = string(secretBuf)
		err := json.Unmarshal([]byte(str), &st1)
		if err != nil {
			return false
		}
	}
	st1 = append(st1, secret)
	buf, _ := json.Marshal(st1)
	if updateSecret(string(buf)){
		return true
	}else {
		return false
	}
}

// 返回secret 列表
func getSecrecList()(secret , error string) {
	fileContent,_ := fileReadByLine(getFilePath())
	if len(fileContent)!=4 {
		return "","文件不合法"
	}
	// 前端的js怎么就不能判断换行符呢，试了好几种办法的都有bug ,最后在这里做处理吧
	tempStr := fileContent[2]
	if tempStr == "\n" {
		tempStr = ""
	}
	return  tempStr,""
}

// md5 加密
func Md5s(s string) string {
	r := md5.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

// sha1 加密
func strToSha1(s string) string {
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

//加密
func enCrypt(orig, key []byte) string {
	//将秘钥中的每个字节累加,通过sum实现orig的加密工作
	sum := 0
	for i := 0; i < len(key); i++ {
		sum += int(key[0])
	}

	//给明文补码
	var pkcs_code = pKCS5Padding(orig, 8)

	//通过秘钥，对补码后的明文进行加密
	for j := 0; j < len(pkcs_code); j++ {
		pkcs_code[j] += byte(sum)
	}
	return base64.StdEncoding.EncodeToString(pkcs_code)
}

//补码
func pKCS5Padding(orig []byte, size int) []byte {
	//计算明文的长度
	length := len(orig)
	padding := size - length % size
	//向byte类型的数组中重复添加padding
	repeats := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(orig, repeats...)
}

//解密
func deCrypt(text string, key []byte) string {
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
	pkcs_unCode := pKCS5UnPadding(orig)
	return string(pkcs_unCode)
}

//去码
func pKCS5UnPadding(orig []byte) []byte {
	//获取最后一位补码的数字
	var tail = int(orig[len(orig) - 1])
	return orig[:(len(orig) - tail)]
}