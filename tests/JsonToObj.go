package main

import (
	"encoding/json"
	"fmt"
)

type SecretTest struct {

	AccountName string
	Password	string
	CreateDate	string
	Describe	string
}

type Transport struct {
	Time  string
	MAC   string
	Id    string
	Rssid string
}

func main() {

	//secrettest_1 := SecretTest{"QQ1","1234567891","shijian1","describe1"}
	//secrettest_2 := SecretTest{"QQ2","1234567892","shijian2","describe2"}
	//var secretSet []SecretTest
	//secretSet = append(secretSet, secrettest_1)
	//secretSet = append(secretSet, secrettest_2)
	//bufs , _ := json.Marshal(secretSet)
	//fmt.Print(string(bufs))
	//
	//var st []Transport
	//t1 := Transport{Time: "22", MAC: "33", Id: "44", Rssid: "55"}
	//st = append(st, t1)
	//t2 := Transport{Time: "66", MAC: "77", Id: "88", Rssid: "99"}
	//st = append(st, t2)
	//buf, _ := json.Marshal(st)
	//fmt.Println(string(buf))
	//
	//
	//
	//var str = string(buf)
	//var st1 []Transport
	//err := json.Unmarshal([]byte(str), &st1)
	//if err != nil {
	//	fmt.Println("some error")
	//}
	//fmt.Println(st1)
	
}
