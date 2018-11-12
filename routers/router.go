package routers

import (
	"github.com/astaxie/beego"
	"secretBox/controllers"
)

func init() {

	//index
	beego.Router("/", &controllers.IndexController{},"*:Index")
	beego.Router("/closepage", &controllers.IndexController{},"post:ClosePage")

	//user
	//beego.Router("/login", &controllers.UserController{},"get:login")
	beego.Router("/user/login", &controllers.UserController{},"post:Login")
	beego.Router("/user/regist", &controllers.UserController{},"get:Reandregist")
	beego.Router("/user/regist", &controllers.UserController{},"post:Regist")


	//secret
	beego.Router("/secret", &controllers.SecretController{},"post:Secret")
	beego.Router("/secret", &controllers.SecretController{},"get:Secretlist")

}