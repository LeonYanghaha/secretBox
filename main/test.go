package main

import "fmt"

//
//import "strconv"
//
//type Person struct {
//	name string
//	age int
//}
//type Students struct {
//	Person
//	class string
//}
//
//func (s Students) study (str,str1 string)  {
//	println(s.name+str+str1+strconv.Itoa(s.age)+s.class)
//}
//
//
//
//func main() {
//	//per := Person{"yu",45}
//   //studeng := Students{per,"math"}
//
//   fd := Students{Person:Person{name:"fg",age:23},class:"df"}
//
//   println(fd.name)
//
//   //studeng.study("34","er")
//
//
//
//}

//func GO() {
//	fmt.Println("我是GO，现在没有发生异常，我是正常执行的。")
//}
//
//func PHP() {
//	panic("我是PHP,我要抛出一个异常了，等下defer会通过recover捕获这个异常，然后正常处理，使后续程序正常运行。")
//	fmt.Println("我是PHP里panic后面要打印出的内容。")
//}
//
//func PYTHON() {
//	fmt.Println("我是PYTHON，没有defer来recover捕获panic的异常，我是不会被正常执行的。")
//}

func GO() {
	fmt.Println("我是GO，现在没有发生异常，我是正常执行的。")
}

func PHP() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("终于捕获到了panic产生的异常：", err) // 这里的err其实就是panic传入的内容
			fmt.Println("我是defer里的匿名函数，我捕获到panic的异常了，我要recover，恢复过来了。")
		}
	}()

	panic("我是PHP,我要抛出一个异常了，等下defer会通过recover捕获这个异常，捕获到我时，在PHP里是不会输出的，会在defer里被捕获输出，然后正常处理，使后续程序正常运行。但是注意的是，在PHP函数里，排在panic后面的代码也不会执行的。")
	fmt.Println("我是PHP里panic后面要打印出的内容。但是我是永远也打印不出来了。因为逻辑并不会恢复到panic那个点去，函数还是会在defer之后返回，也就是说执行到defer后，程序直接返回到main()里，接下来开始执行PYTHON()")
}

func PYTHON() {
	fmt.Println("我是PYTHON，没有defer来recover捕获panic的异常，我是不会被正常执行的。")
}


func main() {
	GO()
	PHP()
	PYTHON()
}
