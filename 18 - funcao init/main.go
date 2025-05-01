package main

import "fmt"

// Define a variável n como uma variável de pacote
var n int

// init é uma função que é chamada automaticamente antes da função main
func init() {
	// Imprime uma mensagem
	println("A função init foi chamada")
	n = 10
}

// A função init é útil para inicializar variáveis ou configurar o ambiente antes que o programa comece a executar
// Ela é chamada automaticamente pelo Go antes da função main
// Você não precisa chamá-la explicitamente
func main() {
	fmt.Println("A função main foi chamada")
	fmt.Println("A variável n é:", n)
}
