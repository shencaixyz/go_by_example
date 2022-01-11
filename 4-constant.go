package main

import (
	"fmt"
	"math"
)

// const 用于声明一个常量
const s string = "我是常量"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
