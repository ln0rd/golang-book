package main

import "fmt"

// o programa morre com o panic
func shouldNotBeOdd(value int) {
	defer recoverExecution()
	if value%2 != 0 {
		panic(fmt.Sprintf("O valor %d é ímpar, e isso não é permitido!", value))
	}
	fmt.Printf("O valor %d é válido (par).\n", value)

}

func recoverExecution() {
	fmt.Println("Recuperando Execução")

	if r := recover(); r != nil {
		fmt.Println("Recuperado Execução, erro: ", r)
		// aqui voce continua a execução do programa
	}
}

// recover consegue capturar o erro e pausar a subida do panic.
func main() {
	shouldNotBeOdd(1)
}
