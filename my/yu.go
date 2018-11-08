package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//func isExist(filePath string) bool {
//	finfo, err := os.Stat(filePath)
//	println(filePath)
//	if err != nil {
//		// no such file or dir
//		return false
//	}
//	if finfo.IsDir() {
//		// it's a file
//		return false
//	} else {
//		// it's a directory
//		return true
//	}
//}

func readFile(path string) {

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		fmt.Print(line)
	}



	//fi, err := os.Open(path)
	//if err != nil {
	//	panic(err)
	//}
	//defer fi.Close()
	//
	//fd, err := ioutil.ReadAll(fi)
	//
	//str := string(fd)
	//
	//return str
}



func main() {
	 readFile("/Users/bigj/.sebox.data")
	//println(isOk)
}
