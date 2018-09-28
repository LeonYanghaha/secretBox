package routers

import (
	"github.com/astaxie/beego"
	"secretBox/controllers"
)

func init() {

	//index
	beego.Router("/", &controllers.IndexController{},"*:Index")

	//user
	//beego.Router("/login", &controllers.UserController{},"get:login")
	beego.Router("/user/login", &controllers.UserController{},"post:Login")


	//secret
	//beego.Router("/save", &controllers.SecretController{},"get:save")
	//beego.Router("/getsecret", &controllers.SecretController{},"get:save")

}