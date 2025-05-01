package main

import "fmt"

func main() {

	// Função anônima é uma função sem nome que é declarada diretamente no código sem a necessidade de ser declarada
	// em um pacote e ser chamada posteriormente. É executada no mesmo momento
	func() {
		fmt.Println("Doing something")
	}()

	func(value string) {
		fmt.Printf("Value %s", value)
	}("Function anonymous")

}
