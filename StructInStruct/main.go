package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `require max: 100`
	Origin string
}

type Bird struct {
	Animal
	Speed  float32
	CanFly bool
}

func main() {
	b := Bird{}

	b.Name = "Emu"
	b.Origin = "Australia"

	b.Speed = 48
	b.CanFly = false

	b2 := Bird{
		Animal: Animal{Name: "Penguin", Origin: "North"},
		Speed:  8,
		CanFly: false,
	}

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag) //>> require max: 100
}
