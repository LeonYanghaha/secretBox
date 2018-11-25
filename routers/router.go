package routers

import (
	"github.com/astaxie/beego"
	"secretBox/controllers"
	"secretBox/toolBox"
)

func init() {

	beego.InsertFilter("/secret/*", beego.BeforeRouter,toolBox.FilterUser)

	//index
	beego.Router("/", &controllers.IndexController{},"*:Index")
	beego.Router("/closepage", &controllers.IndexController{},"post:ClosePage")

	//user
	//beego.Router("/login", &controllers.UserController{},"get:login")
	beego.Router("/user/login", &controllers.UserController{},"post:Login")
	//beego.Router("/user/regist", &controllers.UserController{},"get:Reandregist")
	beego.Router("/user/regist", &controllers.UserController{},"post:Regist")

	//secret
	beego.Router("/secret/addsecret", &controllers.SecretController{},"post:Secret")
	beego.Router("/secret/getsecret", &controllers.SecretController{},"get:Secretlist")

}