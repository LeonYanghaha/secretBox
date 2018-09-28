package controllers

import "github.com/astaxie/beego"

type SecretController struct {
	beego.Controller
}



func (c *SecretController) StaticBlock() {
	println("ig123456789")

	c.Data["Website"] = "gdfb me"
	c.Data["Email"] = "dsfb vgbsfdom"
	c.TplName = "index.tpl"

}