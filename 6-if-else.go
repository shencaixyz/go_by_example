package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 是偶数")
	} else {
		fmt.Println("7 是奇数")
	}

	var a = 7
	if a%4 == 0 {
		fmt.Println(a, "可以被4整除")
	} else {
		fmt.Println(a, "不可以可以被4整除")
	}

	if num := -9; num < 0 {
		fmt.Println(num, "是负数")
	} else if num < 10 {
		fmt.Println(num, "是一位数")
	} else {
		fmt.Println(num, "是两位数以上")
	}
}
