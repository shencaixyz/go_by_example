package main

import "fmt"

func main() {
	//最基础的方式 单个循环条件
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	//经典的初始/条件/后续
	for j := 20; j <= 30; j++ {
		fmt.Println(j)
	}

	//不带条件的 for 循环将一直重复执行
	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
