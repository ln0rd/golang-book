package main

import "fmt"

type Comic struct {
	universe string
}

type Marvel struct {
	Comic
	heros_quantity int32
}

type DC struct {
	Comic
	heros_quantity int32
}

func main() {
	fmt.Println("Inheritance")

	marvel := Marvel{}
	marvel.heros_quantity = 80000
	marvel.universe = "Marvel Multiverse"
	fmt.Println(marvel)

	dc := DC{heros_quantity: 30000, Comic: Comic{universe: "DC Universe"}}
	fmt.Println(dc)

	dc1 := DC{Comic{universe: "Another Universe"}, 1}
	fmt.Println(dc1)
	fmt.Println(dc1.universe)
}
