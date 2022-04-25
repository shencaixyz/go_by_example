package main

import "fmt"

func main() {
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am boole")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Printf("%T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
