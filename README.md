## Personal Golang Handbook

<img src="https://go.dev/images/gophers/motorcycle.svg" width="190" />

This repository contains much content about golang

### Commands, Examples and Guides
Creating a module, this name are going to be the project name
```
go mod init modulename 
```

Creating a module run
```
go mod
```
Removing unecessary dependencies
```
go mod tidy
```
Installing an external package
```
go get github.com/badoux/checkmail
```

Installing your package in your global golang installation, you will can see it installed inside your golang folder installation
```
go install
```

Exemple importing inside main file an submodule
```
package main

import (
	"fmt"
	"mymodule/auxiliar"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")

	auxiliar.Escrever()
}
```

Example mymodule/auxiliar package
```
package auxiliar

import "fmt"

func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	EscreverTwo()//Essa funcao está no mesmo package por isso nao precisa de import
}
```

Adding an external dependencie
```
go get github.com/badoux/checkmail
```
Your go.mod file is going to be like this
```
module email-validator

go 1.19

require github.com/badoux/checkmail v1.2.1 // indirect
```

When you are going to use this, you import using only the last name in import like `import "checkmail"`.