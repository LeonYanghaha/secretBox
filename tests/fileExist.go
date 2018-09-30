package main

import "os"

func main() {

	isExist , _ := PathExists("/home/yu/studygo/src/secretBox/tests/Aw.go")
	println(isExist)

}



func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
