package main

import (
	"fmt"
	"time"
)

func main() {
	privateGetTimeNow()
}

func PublicGetTimeNow() {
	currentTime := time.Now().UTC()
	fmt.Println("Times now: ", currentTime.Format(time.RFC3339))
}

func privateGetTimeNow() {
	currentTime := time.Now().UTC()
	fmt.Println("Times now: ", currentTime.Format(time.RFC3339))
}
