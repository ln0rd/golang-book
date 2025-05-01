package main

import "fmt"

func main() {
	fmt.Println("Switch")

	action := 5
	act := adventurerAction(action)
	fmt.Println(act)

	action = 9
	act = adventurerAction(action)
	fmt.Println(act)

	action = 3
	act = adventurerAction(action)
	fmt.Println(act)

	var itemNameReceived string
	adeventurerBag(2, &itemNameReceived)
	fmt.Println(itemNameReceived)

}

func adventurerAction(number int) string {
	switch number {
	case 1:
		return "Running"
	case 2:
		return "Jumping"
	case 3:
		return "Searching"
	case 4:
		return "Defending"
	case 5:
		return "Spelling Magic"
	default:
		return "Action Not Implemented"
	}
}

func adeventurerBag(item int, itemName *string) {
	switch item {
	case 1:
		*itemName = "Sword"
	case 2:
		*itemName = "Magical Wand"
	case 3:
		*itemName = "Invisibility Helm"
	default:
		*itemName = "Assault Rifle"
	}
}
