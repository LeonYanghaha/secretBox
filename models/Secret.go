package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Secret struct {
	Id          int
	Name        string
	Password	string
	CreateDate	time.Time
	Describe	string
	Uid 		int
}

func init() {

	orm.RegisterModel(new(Secret))

}