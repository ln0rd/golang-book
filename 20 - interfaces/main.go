package main

import "fmt"

type cardTransaction struct {
	ID     int
	amount float64
}

type pixTransaction struct {
	ID     int
	amount float64
	key    string
}

type Transaction interface {
	executeTransaction(t Transaction) int
}

func (ct cardTransaction) executeTransaction(t Transaction) int {
	fmt.Println("Executing card transaction...")
	return ct.ID
}

func (pt pixTransaction) executeTransaction(t Transaction) int {
	fmt.Println("Executing pix transaction...")
	return pt.ID
}

func executeTransaction(t Transaction) int {
	fmt.Println("Starting a new transaction...")
	return t.executeTransaction(t)
}

func main() {
	cardTransaction := cardTransaction{
		ID:     1,
		amount: 100.0,
	}
	pixTransaction := pixTransaction{
		ID:     2,
		amount: 200.0,
		key:    "123456789",
	}

	fmt.Println("Card Transaction ID:", executeTransaction(cardTransaction))
	fmt.Println("Pix Transaction ID:", executeTransaction(pixTransaction))
}
