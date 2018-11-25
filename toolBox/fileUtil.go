package toolBox

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"strconv"
	"time"
)

func updateSecret(line3 string)(isOk bool){

	fileContent,_ := fileReadByLine(GetFilePath())
	if len(fileContent)!=4 {
		return false
	}

	line3 = line3+"\n"
	content:=[]byte(fileContent[0] + fileContent[1] + line3+ fileContent[3])
	err:=writeFile(content,0777)
	if err!=nil {
		fmt.Println(err)
		return false
	}else{
		return true
	}
}

func initFile(un,pw string)(bool,string) {

	//在写入文件之前，应该去检测是否存在历史文件
	hasHistory := HasHistory()
	if hasHistory {
		return false ,"存在历史文件"
	}
	//创建文件
	if !createFile() {
		return false, "文件创建失败"
	}

	line1 := un+"."+pw+"."+ time.Now().Format("2006-01-02 15:04:05")+"\n"
	line2 :=getRandomNum()+"\n"
	line3 :=""+"\n"
	line4 := time.Now().Format("2006-01-02 15:04:05")+"\n"

	content:=[]byte(line1 + line2 + line3 + line4)
	err:=writeFile(content,0777)
	if err!=nil {
		fmt.Println(err)
		return false,"error"
	}else{
		return true ,"sunccess"
	}
}

func writeFile(data []byte,perm os.FileMode) error{

	fileName := GetFilePath()
	f,err:=os.OpenFile(fileName,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,perm)
	if err!=nil {
		return err
	}
	n,err:=f.Write(data)
	if err==nil && n<len(data){
		err=io.ErrShortWrite
	}
	if err1:=f.Close();err==nil{
		err=err1
	}
	return err
}

func createFile()bool{
	userInfoFile := beego.AppConfig.String("userInfoFile")
	file,err:=os.Create(userInfoFile)
	if err!=nil{
		fmt.Println(err)
		return false
	}
	defer file.Close()
	return true
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
	return 99
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