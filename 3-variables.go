package main

import "fmt"

func main() {
	// var 声明1个或者多个变量
	var a = "初始化"
	fmt.Println(a)

	//一次声明多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	//字段推断以及初始化的变量类型
	var d = true
	fmt.Println(d)

	//声明后却没有给出初始值 则变量会初始化未零
	var e int
	fmt.Println(e)

	// := 语法时声明并初始化变量的简写
	f := "short"
	fmt.Println(f)
}
