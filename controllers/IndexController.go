package controllers

import (
	"github.com/astaxie/beego"
	"secretBox/toolBox"
)

type IndexController struct {
	beego.Controller
}



func (c *IndexController) Index() {

	filePath := toolBox.GetFilePath()
	println(filePath)

	isOk := toolBox.CheckFile(filePath)

	type tempData struct {
		isFirst  bool
		isOk     bool
		userName string
	}
	temp := tempData{}

	switch isOk {
		case 0:
			temp.isOk = false
			temp.isFirst = true
			temp.userName = "1234"
			break
		case 1:
			temp.isOk = false
			temp.isFirst = false
			temp.userName = "123"
			break
		default:
			break
	}

	c.Data["isOk"] = true
	c.Data["userName"] = "1234"
	c.Data["isFirst"] = true
	c.TplName = "index.tpl"

}

func (c *IndexController) ClosePage () {


	c.Data["isFirst"] = true
	c.TplName = "index.tpl"

}