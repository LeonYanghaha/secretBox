package main

import (
	"github.com/astaxie/beego"
	_ "secretBox/routers"
)

func main() {

	//check init
	//toolBox.CheckENV()

	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.Run()

}

