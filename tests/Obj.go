package main

import (
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