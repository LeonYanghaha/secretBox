package models

import "github.com/astaxie/beego/orm"

type User struct {
	Uid          int
	Name        string
	Password	string
}

func init() {

	orm.RegisterModel(new(User))

}