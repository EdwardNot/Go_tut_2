package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]
	c := a[3:]
	d := a[:6]
	e := a[3:6]

	fmt.Print(b) // >> [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
	fmt.Print(c) // >> [4, 5, 6, 7, 8, 9, 10]
	fmt.Print(d) // >> [1, 2, 3, 4, 5, 6]
	fmt.Print(e) // >> [4, 5, 6]

	a[5] = 45
	fmt.Print(b) // >> [1, 2, 3, 4, 5, 45, 7, 8, 9, 10]
	fmt.Print(c) // >> [4, 5, 45, 7, 8, 9, 10]
	fmt.Print(d) // >> [1, 2, 3, 4, 5, 45]
	fmt.Print(e) // >> [4, 5, 45]

	maked := make([]int, 3, 100)
	fmt.Println(maked)                     //>> [0, 0, 0]
	fmt.Printf("Length: %v", len(maked))   //>> 3 (count of elements)
	fmt.Printf("Capacity: %v", cap(maked)) //>> 100 (capacity of slice)

	appendable := []int{}
	fmt.Println(appendable)                     //>> []
	fmt.Printf("Length: %v", len(appendable))   //>> 0 (count of elements)
	fmt.Printf("Capacity: %v", cap(appendable)) //>> 0 (capacity of slice)
	appendable = append(appendable, 1)
	fmt.Println(appendable)                     //>> [1]
	fmt.Printf("Length: %v", len(appendable))   //>> 1 (count of elements)
	fmt.Printf("Capacity: %v", cap(appendable)) //>> 2 (capacity of slice)
	//appendable = append(appendable, 2, 3, 4, 5)
	appendable = append(appendable, []int{2, 3, 4, 5}...)
	fmt.Println(appendable)                     //>> [1, 2, 3, 4, 5]
	fmt.Printf("Length: %v", len(appendable))   //>> 5 (count of elements)
	fmt.Printf("Capacity: %v", cap(appendable)) //>> 8 (capacity of slice)

	origin := []int{1, 2, 3, 4, 5}
	newSlice := append(origin[:2], a[3:]...)
	fmt.Println(newSlice) // >> [1, 2, 4, 5]
	fmt.Println(origin)   // >> [1, 2, 4, 5, 5] ???? WHY ????
}
