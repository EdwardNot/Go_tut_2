package main

import (
	"fmt"
)

func main() {
	switch i := 2 + 3; i {
	case 1, 5, 6:
		fmt.Println("One or five or six")
	case 2, 3, 4:
		fmt.Println("Two or three or four")
	default:
		fmt.Println("Another number")
	}

	n := 10
	switch {
	case n <= 10:
		fmt.Println("Less or equal ten")
		fallthrough
	case n <= 20:
		fmt.Println("Less or equal twenty")
	default:
		fmt.Println("Higher than twenty")
	}

	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is integer")
	case float64:
		fmt.Println("i is float")
	case string:
		fmt.Println("i is string")
	default:
		fmt.Println("i is another type")
	}
}
