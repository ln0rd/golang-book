package main

import "fmt"

func main() {
	var result int8 = sum(20, 30)
	println(result)

	var f = func(phrase string) {
		fmt.Println(phrase)
	}

	f("function f")

	var res1, res2, res3 = mathResults(20, 30)
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)

	var _, _, res = mathResults(20, 30)
	fmt.Println(res)
}

func sum(number_one int8, number_two int8) int8 {
	return number_one + number_two
}

func mathResults(n1, n2 int8) (int8, int8, int8) {
	var s int8 = n1 + n2
	var sb int8 = n1 - n2
	var div int8 = n1 / n2
	return s, sb, div
}
