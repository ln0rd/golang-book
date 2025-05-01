package main

import "fmt"

// defer é uma instrução que adia a execução de uma função até o final do bloco atual.
// A função adiada é executada após a função que a contém retornar.

func firstFuncion() {
	fmt.Println("Executando a primeira função")
}

func seconfFunction() {
	fmt.Println("Executando a segunda função")
}

func thirdFunction() {
	fmt.Println("Executando a terceira função")
}

func main() {
	// Ele ira adiar a função ate o ulimo momento possivel
	defer firstFuncion()
	seconfFunction()
	// essa sera execultada primeiro
	defer thirdFunction()
}
