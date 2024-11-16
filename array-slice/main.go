package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Array and Slices")

	var arrayOne [5]int //Array with five positions
	fmt.Println(arrayOne)

	arrayTwo := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arrayTwo)

	arrayTwo[3] = 8
	fmt.Println(arrayTwo)

	var arrayExample [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arrayExample)

	arrayWithLimitInBegin := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(arrayWithLimitInBegin)

	slice := []int{}
	fmt.Println(slice)

	sliceTwo := []int{}
	sliceTwo = append(sliceTwo, 1)
	fmt.Println(sliceTwo)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(arrayOne))

	sliceNew := []int{1, 2, 3, 4, 5}
	sliced := sliceNew[1:2]
	fmt.Println(sliced)

	// internal array
	fmt.Println("-- internal array")
	sliceByMake := make([]float32, 10, 15) // type, size, capacity, tem que ter pelo menos o size
	fmt.Println(sliceByMake)
}
