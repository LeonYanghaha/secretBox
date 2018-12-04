package routers

import (
	"github.com/astaxie/beego"
	"secretBox/controllers"
	"secretBox/filter"
)

func init() {

	beego.InsertFilter("/secret/*", beego.BeforeRouter,filter.FilterUser)
	//index
	beego.Router("/", &controllers.IndexController{},"*:Index")
	beego.Router("/test", &controllers.IndexController{},"get:Test")
	beego.Router("/closepage", &controllers.IndexController{},"get:ExitApp")
	beego.Router("/getfileinfo", &controllers.IndexController{},"get:GetFileInfo")

	//user
	beego.Router("/user/login", &controllers.UserController{},"post:Login")
	beego.Router("/user/regist", &controllers.UserController{},"post:Regist")

	//secret
	beego.Router("/secret/addsecret", &controllers.SecretController{},"post:Secret")
	beego.Router("/secret/getsecret", &controllers.SecretController{},"get:Secretlist")
	beego.Router("/secret/showsecret", &controllers.SecretController{},"post:ShowSecret")
	beego.Router("/secret/delete", &controllers.SecretController{},"post:DeleteSecret")
	beego.Router("/secret/edit", &controllers.SecretController{},"post:EditSecret")
}