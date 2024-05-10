package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("welcome to random number generator")

	rand.Seed(time.Now().UnixNano())
	diesNumber := rand.Intn(6) + 1
	fmt.Println("value of number is ", diesNumber)

	switch diesNumber {
	case 1:
		fmt.Println("yuo can open or move other")
	case 2:
		fmt.Println("move 2 ")
	case 3:
		fmt.Println("move 3")
	case 4:
		fmt.Println("move 4")
	case 5:
		fmt.Println("move 5")
	case 6:
		fmt.Println("move 6 and one more chance")
	default:
		fmt.Println("what heppend ")

	}
}
