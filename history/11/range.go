package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	sum_key := 0
	for key, num := range nums {
		sum += num
		sum_key += key
	}
	fmt.Println("sum:", sum)
	fmt.Println("sum_key:", sum_key)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}
	for i, c := range "goooo" {
		fmt.Println(i, c)
	}
}
