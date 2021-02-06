package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i, j := 0, 0; i < 6; i, j = i+1, j+1 {
		if j == 2 {
			continue
		}
		fmt.Println(i, j)
		if i == 5 {
			break
		}
	}

	s := []int{1, 2, 3}
	for key, value := range s {
		fmt.Println(key, value)
	}

	//while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}
