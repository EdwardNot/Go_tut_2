package main

import "fmt"

func main() {
	if true {
		fmt.Println("This statement is true")
	}

	m := map[string]int{
		"John":  1,
		"Karen": 2,
		"Boris": 3,
	}

	if value, existed := m["John"]; existed {
		fmt.Println(value)
	}

	number := 50
	guess := 30

	if guess < 1 || guess > 100 {
		fmt.Println("Must be between 1 and 100")
	} else {
		fmt.Println("Let's Continue")
	}

	if guess > 1 && guess < 100 {
		if number < guess {
			fmt.Println("Too high")
		} else if number > guess {
			fmt.Println("Too low")
		} else if number == guess {
			fmt.Println("You got it")
		}
	}

}
