package main

import "fmt"

func main() {

	fmt.Println("Welcome to pointer concept")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23

	var ptr = &myNumber

	//fmt.Println("Address of variable : ", &myNumber) // -> & for access address
	//fmt.Println("Value of variable : ", *&myNumber)  // -> *& for access value of that address (*&x and x is same)

	*ptr = *ptr * 2

	fmt.Println("old variable : ", myNumber)

}
