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
	un:=c.GetString("un")
	pw:=c.GetString("pw")
	repw:=c.GetString("repw")
	println(un)
	println(pw)
	println(repw)
	if un=="" || pw=="" || repw==""{
		res.Info ="用户名或者密码不能为空！"
		c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
		c.ServeJSON()
		return
	}

	if pw != repw {
		res.Info ="两次输入的密码不一致！"
		c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
		c.ServeJSON()
		return
	}
	pw  = toolBox.EnCrypetPw(pw)

	isSuccess,info :=toolBox.UserInfoToFile(un,pw)
	println(info)
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

	c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
	c.ServeJSON()
	return
}

func (c *UserController)Reandregist()  {

}

func (c *UserController) Login() {

	name := c.GetString("un")
	pw := c.GetString("pw")
	res := toolBox.GetRes()
	if name=="" || pw==""{
		res.Info ="参数不能为空"
		c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
		c.ServeJSON()
		return
	}

	userinfo,errorinfo := toolBox.GetUserInfo()
	if errorinfo !="" {
		res.Info = "参数不能为空"
		c.Data["json"] = map[string]interface{}{"code": res.Code, "info": res.Info, "data": res.Data}
		c.ServeJSON()
		return
	}
	temp := strings.SplitN(userinfo,".",3)

	if temp[0] !=name {
		res.Info ="用户名错误"
		c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
		c.ServeJSON()
		return
	}

	if temp[1] !=toolBox.EnCrypetPw(pw) {
		res.Info ="密码错误"
		c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
		c.ServeJSON()
		return
	}

	// 成功登录的情况
	timestamp := strconv.FormatInt(time.Now().Unix(),10)
	token := toolBox.GetToken(name, temp[1],timestamp)

	res.Info ="success"
	currentUser := make(map[string]string)
	currentUser["n"]=name
	currentUser["p"]="********"
	currentUser["t"]=token
	currentUser["d"]=timestamp

	res.Code = 1
	res.Data = currentUser
	c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data }
	c.ServeJSON()
	return
}