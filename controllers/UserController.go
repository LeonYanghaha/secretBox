package controllers


import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}



func (c *UserController) Login() {



	name := c.GetString("name")
	pw := c.GetString("pw")

	println(name,pw)

	c.Data["json"] = map[string]interface{}{"success": 0, "message": "111"}
	c.ServeJSON()
	return


}