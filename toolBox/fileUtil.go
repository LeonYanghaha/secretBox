package toolBox

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)


func initFile(un,pw string)bool {
	line1 := un+"."+pw+"."+ time.Now().Format("2006-01-02 15:04:05")
	line2 :=getRandomNum()
	line3 :=""
	line4 := time.Now().Format("2006-01-02 15:04:05")

	flagLine1 := contentToFile(1,line1)
	flagLine2 := contentToFile(2,line2)
	flagLine3 := contentToFile(3,line3)
	flagLine4 := contentToFile(4,line4)
	if flagLine1 && flagLine2 && flagLine3 && flagLine4 {
		return true
	}else{
		return false
	}
}

func contentToFile(lineNum int , content string)bool{



	return false
}

// 0 文件不存在
// 1 文件存在，文件损坏
func CheckFile(filePath string) (int){

	if !fileIsExist(filePath) {
		//如果没有当前文件，肯定是第一次打开，直接返回吧
		return 0
	}
	// 走到这一步，都是文件存在的，然后就校验文件的内容
	lineData,lineNum := fileReadByLine(filePath)

	fmt.Printf("%v",lineData)
	fmt.Printf(strconv.Itoa(lineNum))

	if lineNum != 5{
		return 1
	}

	//user := models.User{}


	return 0

}

func fileReadByLine (path string)([4]string,int){
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	// 用来记录行数
	lineNum := 0
	lineData := [4]string {}
	for {
		lineNum += 1
		if lineNum>5{
			lineNum = 999999
			break
		}
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		fmt.Print(line)
		if lineNum < 5 {
			lineData[lineNum-1] = line
		}
	}
	return lineData,lineNum
}

func fileIsExist(filePath string) (bool){
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		// no such file or dir
		return false
	}
	if fileInfo.IsDir() {
		// it's a directory
		return false
	} else {
		return true
	}
}