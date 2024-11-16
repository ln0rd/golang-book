package main

import (
	"errors"
	"fmt"
)

func main() {
	numbers()
	strings()
	booleans()
	errorsType()

	/*
		Every variable in golang has a initial value
		int = 0
		string = ""
	*/
}

func numbers() {
	// int "architecture based", int8, int16, int32, int64
	var number int16 = 2000
	fmt.Println(number)

	// uint == unsigned int
	// uint, uint8, uint16, uint32, uint64
	var u_number uint32 = 10
	fmt.Println(u_number)

	// rune is equal int32
	var rune_number rune = 30
	fmt.Println(rune_number)

	// byte is equal uint8
	var byte_number byte = 9
	fmt.Println(byte_number)

	// float32, float64
	var float_number float32 = 34.4
	fmt.Println(float_number)

	// float
	float_number_from_architecture := 10.2
	fmt.Println(float_number_from_architecture)
}

func strings() {
	// string
	var str string = "test string"
	fmt.Println(str)

	// char
	char := 'b'
	fmt.Println(char)

	var emptyText string
	fmt.Println(emptyText)
}

func booleans() {
	// true false, value zero is false
	var boolean bool = true
	fmt.Println(boolean)
}

func errorsType() {
	var error_empty error
	fmt.Println(error_empty)

	var error_created error = errors.New("Internal Error")
	fmt.Println(error_created)
}
