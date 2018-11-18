package controllers

import (
	"github.com/astaxie/beego"
	"secretBox/toolBox"
)

type SecretController struct {
	beego.Controller
}
// POST  保存接口
func (c *SecretController)Secret(){

	res := toolBox.GetRes()
	accountName := c.GetString("name")
	password := c.GetString("pw")
	describe := c.GetString("describe")

	if accountName=="" || password==""||describe=="" {
		res.Info ="参数不能为空！"
		c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
		c.ServeJSON()
		return
	}

	res.Info ="123"
	c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
	c.ServeJSON()
	return
}

func (c *SecretController) Secretlist() {

	toolBox.CheckUser()
	res := toolBox.GetRes()
	res.Info ="123"
	c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
	c.ServeJSON()
	return
}
