package filter

import (
	"github.com/astaxie/beego/context"
	"secretBox/toolBox"
	"strconv"
	"time"
)

//实在是找不到可用的文档了，关于beego的资料太少了
// 官网上的文档有报错，导致filter 没法用。
// 最丑的解决办法：自己写个工具类，然后再路由当中调用吧
// 2018-11-18  12:43
//var UpdateCookie = func(c *context.Context){
//
//	n:= c.Input.Cookie("n")
//	newd := strconv.FormatInt(time.Now().Unix(),10)
//	newToken := toolBox.GetToken(n,newd)
//	c.SetCookie("t",newToken)
//	c.SetCookie("d",newd)
//	fmt.Println(newToken,"--","after")
//}

var FilterUser = func(c *context.Context) {

	t:= c.Input.Cookie("t")
	n:= c.Input.Cookie("n")
	d:= c.Input.Cookie("d")
	token := toolBox.GetToken(n,d)

	//fmt.Println(token,"token")
	//fmt.Println(t,"t")
	//fmt.Println(n)
	//fmt.Println(d)

	currentTime := int(time.Now().Unix())
	cookieDate , error := strconv.Atoi(d)
	if error != nil {
		println("cookie 错误-----")
		c.SetCookie("isLogin","false")
		c.Redirect(302, "/")
	}
	//fmt.Println(currentTime)

	if currentTime < cookieDate || currentTime - cookieDate >= 60*10 {
		println("token 过期------")
		c.SetCookie("isLogin","false")
		c.Redirect(302, "/")
	}

	if token != t {
		println("cookie error ------")
		c.SetCookie("isLogin","false")
	    c.Redirect(302, "/")
	}
	//  往下走，就是cookie  合法的，就应该处理业务逻辑，更新cookie
	newd := strconv.FormatInt(time.Now().Unix(),10)
	newToken := toolBox.GetToken(n,newd)
	c.SetCookie("t",newToken)
	c.SetCookie("d",newd)

}
