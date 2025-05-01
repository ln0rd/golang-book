package main

import "fmt"

type Caracter interface {
	doAttack(m Caracter) bool
}

type Hero struct {
	name    string
	lvl     int
	damage  int
	defense int
}

type Monster struct {
	name    string
	lvl     int
	damage  int
	defense int
}

func (h Hero) doAttack(m Caracter) bool {
	// Verifica se o monstro é do tipo Monster
	mt := m.(Monster)

	// Verifica se o ataque do herói é efetivo contra o monstro
	if h.damage > mt.defense {
		fmt.Println("Hero", h.name, "attacked", mt.name, "and it was effective!")
		return true
	}

	fmt.Println("Hero", h.name, "attacked", mt.name, "but it was not effective.")
	return false
}

func (m Monster) doAttack(h Caracter) bool {
	hr := h.(Hero)

	if m.damage > hr.defense {
		fmt.Println("Monster", m.name, "attacked", hr.name, "and it was effective!")
		return true
	}

	fmt.Println("Monster", m.name, "attacked", hr.name, "but it was not effective.")
	return false
}

func main() {
	h := Hero{
		name:    "Bjorn",
		lvl:     1,
		damage:  10,
		defense: 5,
	}

	m := Monster{
		name:    "Dragon",
		lvl:     1,
		damage:  8,
		defense: 3,
	}

	h.doAttack(m)
	m.doAttack(h)
}
