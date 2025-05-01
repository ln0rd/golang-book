package main

import "fmt"

type user struct {
	name string
	age  uint8
}

type company struct {
	id      uint64
	address address
}

type address struct {
	street string
}

func main() {
	fmt.Println("Struc")

	var u user
	u.name = "Dante"
	u.age = 40
	fmt.Println(u)

	u2 := user{}
	u2.name = "Harry"
	u.age = 15
	fmt.Println(u2)

	u3 := user{"Draco", 13}
	fmt.Println(u3)

	u4 := user{name: "Snape"}
	fmt.Println(u4)

	var comp company
	comp.id = 1
	comp.address.street = "address street"
	fmt.Println(comp)

	var comp2 company
	comp2.id = 1
	var add address
	add.street = "street"
	comp2.address = add
	fmt.Println(comp2)
}
