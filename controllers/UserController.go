package controllers

import (
	"github.com/astaxie/beego"
	"secretBox/toolBox"
)

type UserController struct {
	beego.Controller
}



func (c *UserController) Login() {



	name := c.GetString("name")
	pw := c.GetString("pw")
	res := toolBox.GetRes()
	if name=="" || pw==""{
		res.Info ="can shu wei kong !"
		c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }//"success": 0, "message": "111"}
		c.ServeJSON()
		return
	}

	res.Info ="12345"
	res.Code = 0
	c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }//"success": 0, "message": "111"}
	c.ServeJSON()
	return


}