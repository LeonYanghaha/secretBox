package models

import (
	"time"
)

type Secret struct {

	AccountName string
	Password	string
	CreateDate	time.Time
	Describe	string
}