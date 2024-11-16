package main

import "fmt"

func main() {
	// + - / * %
	example_one := 1 + 1
	example_two := 2 - 1
	example_three := 4 / 2
	example_four := 2 * 4
	example_five := 5 % 2

	fmt.Println(example_one, example_two, example_three, example_four, example_five)

	// atribution
	// := =

	// relations operators
	result1 := 1 > 2
	result2 := 1 < 2
	result3 := 1 <= 2
	result4 := 1 >= 2
	result5 := 1 != 2
	result6 := 1 == 1

	fmt.Println(result1, result2, result3, result4, result5, result6)

	// logic
	// && || !
	fmt.Println(false && true)
	fmt.Println(false || true)
	fmt.Println(!false)

	// unarios
	// var++ | var-- | var += var | var -= var | var *= var | var /= var | var %= var
	numero := 10
	numero++
	fmt.Println(numero)
	numero += numero
	fmt.Println(numero)
	numero *= numero
	fmt.Println(numero)
}
