package main

import "fmt"

func main() {

	fmt.Println("Maps")
	adventurer := map[string]string{
		"name": "Adventurer leo",
	}

	fmt.Println(adventurer)
	fmt.Println(adventurer["name"])

	adventurer["level"] = "4"
	fmt.Println(adventurer["level"])

	monster := map[string]map[string]string{
		"atributes": {
			"level": "2",
		},
		"items": {
			"stone": "1",
		},
	}

	fmt.Println(monster)
	fmt.Println(monster["items"]["stone"])
	fmt.Println(monster["atributes"])

	monster["local"] = map[string]string{
		"dungeon": "sewer",
	}
	fmt.Println(monster)

}
