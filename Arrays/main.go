package main

import "fmt"

func main() {
	//var grades [3]int
	//grades := [3]int
	//grades := [3]int{97, 85, 93}
	grades := [...]int{97, 85, 93}

	fmt.Println(grades)

	var students [3]string

	students[0] = "Lisa" //>>[Lisa, , ]

	students[1] = "John"
	students[2] = "Karen"
	fmt.Printf("Number of students: %v\n", len(students))

	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a) //>> 1 2 3
	fmt.Println(b) //>> 1 5 3

	c := [...]int{1, 2, 3}
	d := &a
	d[1] = 5
	fmt.Println(a) //>> 1 5 3
	fmt.Println(b) //>> 1 5 3
}
