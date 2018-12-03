package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"secretBox/models"
	"secretBox/toolBox"
	"strconv"
	"time"
)

type SecretController struct {
	beego.Controller
}

//编辑的功能
func (c *SecretController) EditSecret(){

	res := toolBox.GetRes()
	cSecret := models.Secret{}
	nSecret := models.Secret{}
	cSecret.AppName = c.GetString("current_appname")
	cSecret.Password = c.GetString("current_password")
	cSecret.AccountName = c.GetString("current_accountname")
	cSecret.Describe = c.GetString("current_desc")

	nSecret.AppName = c.GetString("new_appname")
	nSecret.Password = c.GetString("new_password")
	nSecret.AccountName = c.GetString("new_accountname")
	nSecret.Describe = c.GetString("new_desc")
	nSecret.CreateDate = strconv.FormatInt(time.Now().Unix(),10)
	tempRePw := c.GetString("new_repassword")

	if cSecret.AppName==""||cSecret.Password==""||cSecret.AccountName==""||cSecret.Describe==""	 {
		res.Info ="参数不能为空！-c"
		c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
		c.ServeJSON()
		return
	}

	if nSecret.AppName==""||nSecret.Password==""||nSecret.AccountName==""||nSecret.Describe==""||tempRePw==""{
		res.Info ="参数不能为空！-n"
		c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
		c.ServeJSON()
		return
	}
	if nSecret.Password != tempRePw {
		res.Info ="密码不一致"
		c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
		c.ServeJSON()
		return
	}
	// 先检查是否有cSecret 这条数据
	secretStr,error := toolBox.GetSecrecList()
	if error != "" {
		res.Info = error
		c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
		c.ServeJSON()
		return
	}

	secretBuf := []byte(secretStr)
	var str = string(secretBuf)
	var st1 []models.Secret
	err := json.Unmarshal([]byte(str), &st1)
	if err != nil {
		res.Info = err.Error()
		c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
		c.ServeJSON()
		return
	}

	var str2 []models.Secret
	pwSeed := beego.AppConfig.String("pwseedstart")

	var editItem = false
	for i:=0; i < len(st1); i++ {
		signSecret := st1[i]
		//这里的signSecret.password 还需要转换成明文
		tempPassword := toolBox.DeCrypt(signSecret.Password,[]byte(pwSeed))
		if cSecret.AccountName==signSecret.AccountName && cSecret.Describe==signSecret.Describe && cSecret.Password==tempPassword && cSecret.AppName==signSecret.AppName {
			// 查到该item
			// 将nSecret.password 转成密文
			nSecret.Password = toolBox.EnCrypt([] byte(nSecret.Password),[]byte(pwSeed))
			signSecret = nSecret
			editItem = true
		}
		str2 = append(str2,signSecret)
	}
	var isOk bool
	if editItem { // 只有当editItem为true时才写入文件。
		buf, _ := json.Marshal(str2)
		isOk = toolBox.UpdateSecret(string(buf))
	}
	if isOk {
		res.Info = "ok"
		res.Code = 1
	}else{
		res.Info = "文件写入失败！"
		res.Code = -1
	}
	c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
	c.ServeJSON()
	return
}
// 删除
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
			var str2 []models.Secret
			for i:=0; i < len(st1); i++ {
				signSecret := st1[i]
				if accountName==signSecret.AccountName && password==signSecret.Password && appname==signSecret.AppName {
					res.Info = "ok"
					res.Code = 1
				}else{
					str2=append(str2,signSecret)
				}
			}
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
	} else if secretStr=="" {
		res.Code = 1
		res.Info = "ok--当前没有数据的情况"
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
