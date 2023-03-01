package main

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

func main() {
	fmt.Println("Enter Number a:")
	var a int
	var b int
	fmt.Scan(&a)
	fmt.Println("Enter Number b:")
	fmt.Scan(&b)
	fmt.Println(add(a, b))
}
func add(a int, b int) int {
	return a + b
}
