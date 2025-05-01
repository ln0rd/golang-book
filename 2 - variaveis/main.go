package main

import (
	"fmt"
	"strconv"
)

func main() {

	// explicit
	var value string = "first variable"
	fmt.Println(value)

	// inference
	secondValue := "second variable"
	fmt.Println(secondValue)

	var (
		thirdValue   string = "thirdValue"
		anotherValue string = "anotherValue"
	)
	fmt.Println(thirdValue + " " + anotherValue)

	name, birth := "leo", 1996
	birthString := strconv.Itoa(birth)
	fmt.Println("name " + name + " birth " + birthString)

	// change values
	thirdValue, anotherValue = anotherValue, thirdValue
}
