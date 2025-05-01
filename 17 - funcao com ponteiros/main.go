package main

import "fmt"

func inverteSinal(number int) int {
	return number * -1
}

func inverteSinalByPointerReference(number *int) {
	*number = *number * -1
}

func main() {
	number := 10
	result := inverteSinal(number)

	fmt.Println("O número invertido é:", result)

	numberPointer := 11
	inverteSinalByPointerReference(&numberPointer)
	fmt.Println("O número invertido por referência é:", numberPointer)
}
