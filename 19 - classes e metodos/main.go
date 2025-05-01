package main

import (
	"fmt"
	"time"
)

type transaction struct {
	ID             int
	amount         float64
	settlementDate time.Time
}

func (t *transaction) antecipateToday() time.Time {
	t.settlementDate = time.Now()
	return t.settlementDate
}

func (t *transaction) setID(id int) {
	t.ID = id
}

func (t transaction) toString() string {
	return fmt.Sprintf("ID: %d, Amount: %.2f, Settlement Date: %s", t.ID, t.amount, t.settlementDate.Format("2006-01-02"))
}

func main() {
	t := transaction{
		ID:             1,
		amount:         100.0,
		settlementDate: time.Now().Add(30 * 24 * time.Hour),
	}

	fmt.Println(t.toString())

	t.setID(2)
	t.antecipateToday()
	fmt.Println(t.toString())
}
