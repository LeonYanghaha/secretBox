package toolBox

func CheckENV()  {
//	userInfoFile := beego.AppConfig.String("userInfoFile")
//	homePath, _ :=  getUserHome()
//	filePath := path.Join(homePath,userInfoFile)
//	println(filePath)
//
//	isOk := checkFile(filePath)
//	app := models.GetApp()
//
//	switch isOk {
//
//		case 0:
//			app.IsFirst = true
//			app.Status = false
//			app.Username = "00000"
//			break
//		case 1:
//			app.IsFirst = false
//			app.Status = false
//			app.Username = "11111"
//			break
//		default:
//			break
//
//	}
}

// 0  first
// 1  already regist
//func checkIsFirst(filePath string) string {
//
//	isExist,_ := checkFile(filePath)
//	if !isExist {
//		return strconv.Itoa(0)
//	}
//	//文件存在，去解析文件的内容、
//	 // strFileContent,_ := readFile(filePath)
//
//
//	return strconv.Itoa(0)
//}

//func readFile(filePath string)  {
//
//	f, err := os.Open(filePath)
//
//	var str [2] string
//	if err != nil {
//		return nil, err
//	}
//	buf := bufio.NewReader(f)
//	for {
//		line, err := buf.ReadString('\n')
//		line = strings.TrimSpace(line)
//
//		//handler(line)
//
//		if err != nil {
//			if err == io.EOF {
//				return nil
//			}
//			return "", err
//		}
//	}
//	return nil
//}

//func checkFile(path string)( bool,error) {
//	_, err := os.Stat(path)
//	if err == nil {
//		return true, nil
//	}
//	if os.IsNotExist(err) {
//		return false, nil
//	}
//	return false, err
//
//}


//func OpenUrl(){
//
//	url := "http://127.0.0.1:"+beego.AppConfig.String("httpport")
//	println(url)
//	cmd := exec.Command("explorer", url)
//	err := cmd.Start()
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//}
