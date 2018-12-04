package toolBox

import (
	"strings"
)

// 加密当前用户的密码
func EnCrypetPw(pw string)string{
	tempStr := "9olmjuil+_{"+pw+"|}{POKM<MNBGFTJ"
	return strToSha1(tempStr)
}

func getPassStr()(passStr,error string){
	userinfo, errorinfo := getUserInfo()
	if errorinfo!="" {
		return "" ,errorinfo
	}
	return (strings.SplitN(userinfo,".",3))[1],""
}

// 传入用户的信息，加密，生成token
func getToken(un,timestamp string) string {
	temp := GetTokenSeed() //beego.AppConfig.String("tokenseed")
	return strToSha1(un + temp + timestamp)
}
// 返回当前的用户信息
func getUserInfo ()(userinfo,errorinfo string){
	fileContent,_ := fileReadByLine(getFilePath())
	if len(fileContent)!=4 {
		return "","文件不合法"
	}
	return fileContent[0],""
}