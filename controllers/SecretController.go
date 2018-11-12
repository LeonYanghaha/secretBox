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


	println(c.GetSession("pw"))
	println(c.GetSession("name"))

	s := c.GetSession("name")
	println(&s)



	res := toolBox.GetRes()

	accountName := c.GetString("name")
	password := c.GetString("pw")
	describe := c.GetString("describe")
	//createDate :=  time.Now().Format("2006-01-02 15:04:05")

	if accountName=="" || password==""||describe=="" {
		res.Info ="参数不能为空！"
		c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
		c.ServeJSON()
		return
	}

	//secret := new models.Secret{}


	res.Info ="123"
	c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
	c.ServeJSON()
	return



}

func (c *SecretController) Secretlist() {


}
