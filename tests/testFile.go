package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	//fileInfo, _ := os.Stat("/home/yu/studygo/src/secretBox/tests/default_test.go")
	//fmt.Println("FIle name:", fileInfo.Name())
	//fmt.Println("Size in bytes:", fileInfo.Size())
	//fmt.Println("Permissions:", fileInfo.Mode())
	//fmt.Println("Last modified:", fileInfo.ModTime())
	//fmt.Println("Is directory:", fileInfo.IsDir())
	//fmt.Printf("System interface type:%T\n", fileInfo.Sys())
	//fmt.Printf("System info:%+v\n\n", fileInfo.Sys())


	//
	//user, _ := user.Current()
	//
	//println(user.HomeDir)


	ReadLine("/home/yu/studygo/src/secretBox/tests/default_test.go", Print)

}



func ReadLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}
func Print(line string) {
	fmt.Println(line)
}

