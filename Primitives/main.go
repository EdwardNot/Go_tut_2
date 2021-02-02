package main

import "fmt"

var n bool = true // or false || default value for bool is "false"

func main() {
	m := 1 == 1

	fmt.Println(m)

	a := 10 // 1010
	b := 3  // 0011

	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 1011
	fmt.Println(a ^ b)  // 1001
	fmt.Println(a &^ b) // 0100

	a = 8

	// Move bytes
	fmt.Println(a << 3)
	fmt.Println(a >> 3)
}
