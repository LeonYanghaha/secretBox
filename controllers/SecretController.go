package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"secretBox/models"
	"secretBox/toolBox"
	"strconv"
	"time"
)

type SecretController struct {
	beego.Controller
}

func (c *SecretController) DeleteSecret(){
	res := toolBox.GetRes()
	accountName := c.GetString("ac")
	password := c.GetString("pw")
	appname := c.GetString("an")
	if accountName=="" || password==""||appname=="" {
		res.Info ="参数不能为空！"
		c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
		c.ServeJSON()
		return
	}

	secretStr,err := toolBox.GetSecrecList()
	if err != "" {
		res.Info = err
	}else {
		secretBuf := []byte(secretStr)
		var str = string(secretBuf)
		var st1 []models.Secret
		err := json.Unmarshal([]byte(str), &st1)
		if err != nil {
			res.Info = "error"
			res.Code = -1
		}else{
			//secretList  := make(map[string] []models.Secret)
			//secretList["secret"] = st1
			//res.SecretData = secretList
			fmt.Print(st1)
			var str2 []models.Secret
			for i:=0; i < len(st1); i++ {
				signSecret := st1[i]
				println(signSecret.AccountName)
				if accountName==signSecret.AccountName && password==signSecret.Password && appname==signSecret.AppName {
					res.Info = "ok"
					res.Code = 1
				}else{
					str2=append(str2,signSecret)
				}
			}


			//secretBuf := []byte(secretStr)
			//
			//var str = string(secretBuf)
			//var st1 []models.Secret
			//err := json.Unmarshal([]byte(str), &st1)
			//if err != nil {
			//	return false
			//}
			//
			//st1 = append(st1, secret)
			//buf, _ := json.Marshal(st1)
			////fmt.Println(string(buf))
			//if updateSecret(string(buf)){
			//	return true
			//}else {
			//	return false
			//}

			buf, _ := json.Marshal(str2)
			isOk := toolBox.UpdateSecret(string(buf))
			if isOk {
				res.Info = "ok"
				res.Code = 1
			}else{
				res.Info = "文件写入失败！"
				res.Code = -1
			}
		}
	}
	c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.SecretData }
	c.ServeJSON()
	return
}

// 展示密码的明文
func (c *SecretController) ShowSecret(){
	res := toolBox.GetRes()
	secret := c.GetString("secret")
	if secret=="" {
		res.Info ="参数不能为空！"
		c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data }
		c.ServeJSON()
		return
	}
	pwSeed := beego.AppConfig.String("pwseedstart")
	realSecret := toolBox.DeCrypt(secret,[]byte(pwSeed))
	secretList := make(map[string] string)
	secretList["secret"] = realSecret
	res.Code = 1
	res.Info = "success"
	res.Data = secretList
	c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data }
	c.ServeJSON()
	return
}


// POST  保存接口
func (c *SecretController)Secret(){

	res := toolBox.GetRes()
	accountName := c.GetString("ac")
	password := c.GetString("pw")
	repassword := c.GetString("repw")
	describe := c.GetString("desc")
	appname := c.GetString("an")
	if accountName=="" || password==""||describe==""||appname==""||repassword=="" {
		res.Info ="参数不能为空！"
		c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
		c.ServeJSON()
		return
	}

	if repassword!= password {
		res.Info ="两次输入的密码不一致"
		c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
		c.ServeJSON()
		return
	}

	currentTime := strconv.FormatInt(time.Now().Unix(),10)
	pwSeed := beego.AppConfig.String("pwseedstart")
	secretPW := toolBox.EnCrypt([]byte(password),[]byte(pwSeed))
	secret := models.Secret{appname,accountName,secretPW,currentTime,describe}
	if  toolBox.SaveSecretToFile(secret) {
		res.Info = "ok"
		res.Code = 0
	}else {
		res.Info = "error"
	}
	c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
	c.ServeJSON()
	return
}

func (c *SecretController) Secretlist() {

	res := toolBox.GetRes()
	secretStr,err := toolBox.GetSecrecList()
	if err != "" {
		res.Info = err
	}else {
		secretBuf := []byte(secretStr)
		var str = string(secretBuf)
		var st1 []models.Secret
		err := json.Unmarshal([]byte(str), &st1)
		if err != nil {
			res.Info = "error"
			res.Code = -1
		}else{
			res.Info = "ok"
			res.Code = 1
			secretList  := make(map[string] []models.Secret)
			secretList["secret"] = st1
			res.SecretData = secretList
		}
	}
	c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.SecretData    }
	c.ServeJSON()
	return
}
