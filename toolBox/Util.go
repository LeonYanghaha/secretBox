package toolBox

import (
	"math/rand"
	"secretBox/models"
	"strconv"
)
//  对各个工具包封装，对外提供接口调用。
// 各个工具类不能被直接调用

func GetTokenSeed()string{
	// 这里原来是打算用主机名的，但是主机名可能会发生变化，
	// 思来想去，写死的一个字符串吧。
	return 	"yuik,mnjhiol"+GetPassSeed()+")P:>L<MKLOP"
}

func GetPassSeed()string{
	pass, error := getPassStr()
	if error !="" {
		return ""
	}else{
		return pass
	}
}

func GetUserInfo()(userinfo , error string){
	return getUserInfo()
}

func GetToken(un,timestamp string) string{
	return getToken(un,timestamp)
}

func SaveSecretToFile(secret models.Secret)bool{
	return saveSecret(secret)
}

func EnCrypt(str string)string{
	return enCrypt([]byte(str),[]byte(GetPassSeed()))
}

func DeCrypt(str string) string{
	return deCrypt(str,[]byte(GetPassSeed()))
}

func GetSecrecList()(secretList,error string) {
	return getSecrecList()
}

func GetFilePath()string{
	return getFilePath()
}

func UpdateSecret(str string )(isOk bool){
	return updateSecret(str)
}

func HasHistory()bool{
	return fileIsExist(getFilePath())
}

func UserInfoToFile(un,pw string) (bool,string) {
	return initFile(un,pw)
}

func getRandomNum ()string{
	num := RandInt64()
	str := strconv.FormatInt(num,10)
	return strToSha1(str)
}

func GetRes() models.Res {
	var res models.Res
	res.Code = -1
	res.Info = "error"
	return res
}

func RandInt64() int64 {
	return rand.Int63n(999999-100000) + 100000
}