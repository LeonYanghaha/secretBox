package main

import (
	"github.com/astaxie/beego"
	_ "secretBox/routers"
	"secretBox/toolBox"
)

func main() {

	//check init
	toolBox.CheckENV()
	beego.Run()
}

