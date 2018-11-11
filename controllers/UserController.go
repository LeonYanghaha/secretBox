package controllers

import (
	"github.com/astaxie/beego"
	"secretBox/toolBox"
)

type UserController struct {
	beego.Controller
}
// POST  注册接口
func (c *UserController)Regist()  {

	res := toolBox.GetRes()
	un:=c.GetString("username")
	pw:=c.GetString("password")
	repw:=c.GetString("repassword")
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
	
	name := c.GetString("name")
	pw := c.GetString("pw")
	res := toolBox.GetRes()
	if name=="" || pw==""{
		res.Info ="can shu wei kong !"
		c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }//"success": 0, "message": "111"}
		c.ServeJSON()
		return
	}

	res.Info ="12345"
	res.Code = 0
	c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }//"success": 0, "message": "111"}
	c.ServeJSON()
	return

}