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
	res := toolBox.GetRes()
	isOk := toolBox.CheckFile(filePath)

	switch isOk {
	case 0 :
		res.Code = -1
		res.Info = "貌似是第一次登录，请先去注册吧."
		break
	case 1 :
		res.Code = -1
		res.Info= "文件被损坏了."
		break
	case 99 :
		res.Code = 1
		res.Info= "ok"
		break
	default:
		res.Code = -1
		res.Info= "这是不可能出现的"
		break
	}
	c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data}
	c.ServeJSON()
	return
}

func (c *IndexController) ClosePage () {


	c.Data["isFirst"] = true
	c.TplName = "index.tpl"

}