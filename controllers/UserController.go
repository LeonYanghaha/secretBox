package controllers

import (
	"github.com/astaxie/beego"
	"secretBox/toolBox"
	"strings"
)

type UserController struct {
	beego.Controller
}

//var globalSessions *session.Manager
//
//func init() {
//	sessionConfig := &session.ManagerConfig{
//		CookieName:beego.AppConfig.String("appname"),
//		EnableSetCookie: true,
//		Gclifetime:3600,
//		Maxlifetime: 3600,
//		Secure: false,
//		CookieLifeTime: 3600,
//		ProviderConfig: "./tmp",
//	}
//	globalSessions, _ = session.NewManager("memory",sessionConfig)
//	go globalSessions.GC()
//}

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

	//sess, _ := globalSessions.SessionStart(w, r)
	//defer sess.SessionRelease(w)
	//username := sess.Get("username")
	//if r.Method == "GET" {
	//	t, _ := template.ParseFiles("login.gtpl")
	//	t.Execute(w, nil)
	//} else {
	//	sess.Set("username", r.Form["username"])
	//}

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

	//println(ctx.ResponseWriter, c.Request)
	//sess, _ := globalSessions.SessionStart(w, r)
	//defer sess.SessionRelease(w)

	c.SetSession("name",name)
	c.SetSession("pw",pw)



	//sess.Set("name",name)
	//sess.Set("pw", pw)

	res.Info ="success"
	currentUser := make(map[string]string)
	currentUser["name"]=name
	currentUser["pw"]="*************"
	res.Data = currentUser
	c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
	c.ServeJSON()
	return
}