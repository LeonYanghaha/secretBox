package toolBox

import (
	"fmt"
	"github.com/astaxie/beego/context"

)

//实在是找不到可用的文档了，关于beego的资料太少了
// 官网上的文档有报错，导致filter 没法用。
// 最丑的解决办法：自己写个工具类，然后再路由当中调用吧
// 2018-11-18  12:43

//type Filer struct {
//	beego.Controller
//}
//
//func (c *Filer) UserFiler(){
//	println("-------------filter--------------")
//}

var FilterUser = func(c *context.Context) {
	t:= c.Input.Cookie("t")
	//p:= c.Input.Cookie("p")
	n:= c.Input.Cookie("n")
	d:= c.Input.Cookie("d")
	token := GetToken(n,d)

	//TODO：判断是否过期
	fmt.Println(token)
	fmt.Println(t)
	//fmt.Println(p)
	fmt.Println(n)
	fmt.Println(d)
	if token != t {
		println("cookie error ------")
		c.Redirect(302, "/")
	}
}
