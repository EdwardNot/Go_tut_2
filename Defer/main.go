package main

import "fmt"

func main() {
	fmt.Println("start")
	fmt.Println("middle")
	fmt.Println("end")

	//* >> start \n middle \n end

	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")

	//* >> start \n end \n middle

	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")

	//* >> end \n middle \n start

	//* defer откладывает указанную комманду до конца выполнения текущего кода
	//* таким образом можно открывать и закрывать запросы в одном месте кода
	//* {
	//* 	открыть запрос
	//* 	defer закрыть запрос
	//* 	операция с запросом
	//* } запрос закроектся только после операций

	a := "start"
	defer fmt.Println(a)
	a = "end"

	//* будет выведено "start", так как defer принимает в себе аргументы на момент объявления

}
