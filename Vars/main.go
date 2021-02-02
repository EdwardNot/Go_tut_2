package main

import (
	"fmt"
	"strconv"
)

var fl float32 = 25.6 // global variable

var (
	name    string = "Jhon Doe"
	company string = "Some Company"
	id      int    = 95589
)

const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

const (
	d = iota
	e
	f
)

const (
	h = iota + 5
	j
	k
)

const (
	_  = iota //ignore 0 value
	KB = 1 << (iota * 10)
	MB
	GB
	TB
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	var i int = 34 // equals i := 34
	i = 42

	fmt.Println(i)

	var j int
	k := 34

	fmt.Printf("%v, %T", k, j) // %v - value; %T - Type

	fl = float32(i) //cast to type

	var intStrint string = "123"

	intStrint = strconv.Itoa(i)

	fmt.Println(intStrint)

	fileSize := 40000000000
	fmt.Printf("%.2fGB", fileSize/GB)

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v", isHeadquarters&roles == isHeadquarters)
}
