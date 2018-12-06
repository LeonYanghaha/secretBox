package main

import (
	"github.com/astaxie/beego"
	"runtime"
	_ "secretBox/routers"
)


func getCurrentPath() string {

	_, file, _, ok := runtime.Caller(1)
	if !ok {
		//panic(errors.New("Can not get current file info"))
		return ""
	}
	return file

	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	beego.Debug(err)
	//}
	//return strings.Replace(dir, "\\", "/", -1)
}

func main() {

	//beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	//	AllowAllOrigins:  true,
	//	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	//	ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	//	AllowCredentials: true,
	//}))
	//fmt.Println(getCurrentPath(),"--")
	//println(getCurrentPath())
	beego.SetViewsPath("views")
	beego.SetStaticPath("/static","./views/static/")
	beego.Run()
}

