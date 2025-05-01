package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		fmt.Println("I Value: ", i)
		time.Sleep(time.Second / 20)
		i++
	}

	for j := 0; j < 10; j++ {
		fmt.Println("J value: ", j)
		time.Sleep(time.Second / 60)
	}

	names := [3]string{"Sasuke", "Naruto", "Sakura"}
	// for indice, name := range array
	for index, name := range names {
		fmt.Printf("index [ %d ] value [ %s ] \n", index, name)
	}

	for index, letter := range "ROHAN" {
		fmt.Printf("Index: %d Letter: %s \n", index, string(letter))
	}

	user := map[string]string{
		"name":  "leo",
		"type":  "adventurer",
		"class": "mage",
	}
	for key, value := range user {
		fmt.Printf("{ %s } { %s } \n", key, value)
	}

}
