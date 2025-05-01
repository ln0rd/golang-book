package main

import "fmt"

func fibbonacci(position uint) uint {
	if position <= 1 {
		return uint(position)
	}

	return fibbonacci(position-2) + fibbonacci(position-1)
}

func main() {
	// Função recursiva
	result := fibbonacci(10)
	fmt.Println(result)
}
