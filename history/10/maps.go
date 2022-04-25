package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	fmt.Println("k1=", m["k100"])

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("new map:", m)

	_, prs := m["k100"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
