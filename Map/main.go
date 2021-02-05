package main

import "fmt"

func main() {
	ids := map[string]int{
		"John":  1,
		"Karen": 2,
		"Boris": 3,
	}

	fmt.Println(ids)

	mapWithArrayKey := map[[3]int]string{} //Only Array can be key, not slice

	fmt.Println(mapWithArrayKey)

	mapMaked := make(map[string]int) // mapMaked = make(map[string]int, 10)
	mapMaked = map[string]int{
		"John":  1,
		"Karen": 2,
		"Boris": 3,
	}
	mapMaked["George"] = 4
	fmt.Println(mapMaked["George"])
	delete(mapMaked, "George")

	value, existance := mapMaked["Jhon"]
	fmt.Println(value, existance) // >> 1, true

	value, existance = mapMaked["Ben"]
	fmt.Println(value, existance) // >> 0, false

	//* you ca write _, existance := mapMaked["Jhon"] if value is not important

}
