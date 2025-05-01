package main

import "fmt"

func calculoComRetornoNomeado(n1, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	soma, sub := calculoComRetornoNomeado(4, 2)
	fmt.Println(soma, sub)
	fmt.Println(calculoComRetornoNomeado(4, 3))
}
