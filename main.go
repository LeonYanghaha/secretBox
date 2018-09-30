package main

import (
	_ "secretBox/routers"
	"github.com/astaxie/beego"
	"secretBox/toolBox"
)

func main() {

	//check init
	toolBox.Init()
	toolBox.OpenUrl()
	beego.Run()


}

