package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {

	fmt.Println("Main package")
	erro := checkmail.ValidateFormat("123")
	fmt.Println(erro)

	if erro != nil {
		fmt.Println("Erro no formato do email")
	}
}
