package main

import (
	"fmt"
	"log"
)

func main() {
	callPanic()
	fmt.Println("Мы восстановились после паники!")
}
func callPanic() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Случилась паника:", err)
		}
	}()
	fmt.Println(selectFromList())
}
func selectFromList() string {
	list := []string{
		"first",
		"second",
		"third",
	}
	// обращаемся к индексу, которого нет в массиве, тем самым, вызываю неявную панику
	return list[3]
}
