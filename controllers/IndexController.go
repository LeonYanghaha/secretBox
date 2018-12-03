package controllers

import (
	"github.com/astaxie/beego"
	"secretBox/toolBox"
)

type IndexController struct {
	beego.Controller
}

//  返回文件信息
func (c *IndexController) GetFileInfo () {

	filePath := toolBox.GetFilePath()
	res := toolBox.GetRes()
	fileInfo := make(map[string]string)
	fileInfo["fileInfo"]=filePath
	res.Data = fileInfo
	res.Info = "success"
	res.Code = 1
	c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data}
	c.ServeJSON()
	return
}


// 首页的打开，用于检测当前状态
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
		res.Info= "一切ok，可以直接登录👌"
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
// 退出时，关闭各种资源。
func (c *IndexController) ClosePage () {


	c.Data["isFirst"] = true
	c.TplName = "index.tpl"

}