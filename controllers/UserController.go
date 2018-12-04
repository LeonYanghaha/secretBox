package controllers

import (
	"github.com/astaxie/beego"
	"secretBox/toolBox"
	"strconv"
	"strings"
	"time"
)

type UserController struct {
	beego.Controller
}

// POST  注册接口
func (c *UserController)Regist()  {

	res := toolBox.GetRes()
	defer func() {
		c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
		c.ServeJSON()
		return
	}()
	un:=c.GetString("un")
	pw:=c.GetString("pw")
	repw:=c.GetString("repw")
	if un=="" || pw=="" || repw==""{
		res.Info ="用户名或者密码不能为空！"
		return
	}

	if pw != repw {
		res.Info ="两次输入的密码不一致！"
		return
	}
	pw  = toolBox.EnCrypetPw(pw)

	isSuccess,_ :=toolBox.UserInfoToFile(un,pw)
	if isSuccess{
		res.Info ="注册成功！"
		res.Code=1
		userInfo := make(map[string]string)
		userInfo["un"]=un
		userInfo["pw"]=pw
		res.Data=userInfo
	}else{
		res.Info ="文件写入失败！"
	}
	return
}

func (c *UserController) Login() {

	name := c.GetString("un")
	pw := c.GetString("pw")
	res := toolBox.GetRes()
	defer func() {
		c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
		c.ServeJSON()
		return
	}()
	if name=="" || pw==""{
		res.Info ="参数不能为空"
		return
	}
	userinfo,errorinfo := toolBox.GetUserInfo()
	if errorinfo !="" {
		res.Info = "参数不能为空"
		return
	}
	temp := strings.SplitN(userinfo,".",3)

	if temp[0] !=name {
		res.Info ="用户名错误"
		return
	}

	if temp[1] !=toolBox.EnCrypetPw(pw) {
		res.Info ="密码错误"
		return
	}

	// 成功登录的情况
	timestamp := strconv.FormatInt(time.Now().Unix(),10)
	token := toolBox.GetToken(name,timestamp)
	res.Info ="success"
	currentUser := make(map[string]string)
	currentUser["n"]=name
	currentUser["p"]="********"
	currentUser["t"]=token
	currentUser["d"]=timestamp
	res.Code = 1
	res.Data = currentUser
	return
}