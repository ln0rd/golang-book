package main

import "fmt"

func main() {
	fmt.Println("Estrutura de Controle")

	number := 10

	// if else
	if number > 9 {
		fmt.Println("bigger than")
	} else {
		fmt.Println("not bigger than")
	}

	// if init
	if anotherNumber := number; anotherNumber > 0 {
		fmt.Println("bigger than")
	}

}
