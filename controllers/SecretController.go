package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"secretBox/models"
	"secretBox/toolBox"
	"strconv"
	"time"
)

type SecretController struct {
	beego.Controller
}
// POST  保存接口
func (c *SecretController)Secret(){

	res := toolBox.GetRes()
	accountName := c.GetString("ac")
	password := c.GetString("pw")
	describe := c.GetString("desc")

	if accountName=="" || password==""||describe=="" {
		res.Info ="参数不能为空！"
		c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
		c.ServeJSON()
		return
	}

	currentTime := strconv.FormatInt(time.Now().Unix(),10)
	fmt.Print(currentTime)

	secret := models.Secret{accountName,password,currentTime,describe}
	if  toolBox.SaveSecretToFile(secret) {
		res.Info = "ok"
		res.Code = 0
	}else {
		res.Info = "error"
	}
	c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
	c.ServeJSON()
	return
}

func (c *SecretController) Secretlist() {

	res := toolBox.GetRes()
	secretStr,err := toolBox.GetSecrecList()
	if err != "" {
		res.Info = err
	}else {
		res.Info = "ok"
		res.Code = 0
		secretList  := make(map[string]string)
		secretList["secretStr"] = secretStr
		res.Data = secretList
	}
	c.Data["json"] = map[string]interface{}{ "code":res.Code , "info":res.Info,"data":res.Data    }
	c.ServeJSON()
	return
}
