package main

import "fmt"

func main() {
	fmt.Println("start")
	panic("Something heppend")
	fmt.Println("end")

	// >> start \n ERROR: Something heppend
	//* позволяет завершить исполнение кода с ошибкой
	//* defer позволит исполнить код до конца, а потом выкинет ошибку

	fmt.Println("start")
	defer panic("Something heppend")
	fmt.Println("end")
}
