package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Friday?")
	today := time.Now().Weekday()
	fmt.Println("Today is", today)
	fmt.Println("Tomorrow is", today+1)
	switch time.Friday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
