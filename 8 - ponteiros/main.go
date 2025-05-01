package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Ponteiros")

	// In this case Im passing the value by copy
	var exp_var int8 = 10
	var exp_var_pointer int8 = exp_var
	fmt.Println(exp_var, exp_var_pointer)
	exp_var_pointer++
	fmt.Println(exp_var, exp_var_pointer)

	// Using pointer
	var exp_var_pt int8 = 8
	var pointer_to_exp_var_pt *int8 = &exp_var_pt
	fmt.Println(exp_var_pt, pointer_to_exp_var_pt)

	*pointer_to_exp_var_pt++ // updating exp_var_pt
	fmt.Println(exp_var_pt, pointer_to_exp_var_pt)
	fmt.Println(exp_var_pt, *pointer_to_exp_var_pt)

	var value int8 = 12
	execFirstExample(&value) // need use & to pass the value in params
	valueInString := strconv.Itoa(int(value))
	fmt.Println("Value after changed by pointer " + valueInString)
}

func execFirstExample(value_pointer *int8) { // it will receive a pointer using this syntax * before type
	fmt.Println(value_pointer)
	*value_pointer++
}
