package main

import "fmt"

type Doctor struct {
	number     int
	actorName  string
	companions []string
	Episodes   []string //public field
}

func main() {

	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	bDoctor := Doctor{3, "Jon Pertwee", []string{"Liz Shaw", "Jo Grant", "Sarah Jane Smith"}}

	fmt.Println(aDoctor, bDoctor)
	fmt.Println(aDoctor.companions[1])

	simpleStruct := struct{ value string }{value: "this is value"}
	fmt.Println(simpleStruct)

}
