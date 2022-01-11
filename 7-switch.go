package main

import (
	"fmt"
	"time"
)

func main() {
	//一个基础的 switch
	i := 3
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// 在同一个 case 预计里面 可以用逗号来分割多个表达式
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println(time.Now().Weekday(), "是周末")
	default:
		fmt.Println(time.Now().Weekday(), "是工作日")
	}

	//不带表达式的 switch 是实现if/else 逻辑的另外一种方式
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println(t.Hour(), "是上午")
	default:
		fmt.Println(t.Hour(), "是下午")
	}

	//类型开关 tyep switch 比较类型而非值 可以用来发现一个接口值得类型
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("布尔类型")
		case int:
			fmt.Println("整型")
		default:
			fmt.Println("未知类型", t)
		}

	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
