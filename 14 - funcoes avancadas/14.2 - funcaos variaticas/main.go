package main

import "fmt"

func somaComVariatica(numbers ...int) (total int) {
	total = 0

	for _, value := range numbers {
		total += value
	}

	return
}

func main() {
	fmt.Println(somaComVariatica(4, 3, 2, 6))
}
