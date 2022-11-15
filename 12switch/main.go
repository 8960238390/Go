package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("switch example")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("Dice value id 1 you can open")
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("you can move to 3 spot")
	case 4:
		fmt.Println("you can move to 4 spot")
	case 5:
		fmt.Println("you can move 5 sport")
	case 6:
		fmt.Println("you can move 6 spot and roll a dice again")
	default:
		fmt.Println("what was that")
	}

}
