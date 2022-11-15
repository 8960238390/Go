package main

import "fmt"

func main() {

	fmt.Println("function Example")

	greet()

	// greet1 := greet
	// greet1()

	result := adder(2, 5)

	fmt.Println("Addition is : ", result)

	proResult, message := proAdder(2, 3, 4, 5)
	fmt.Println("Pro Result is ", proResult, message)
}

func greet() {
	fmt.Println("Welcome to greet GO")
}

func adder(num1 int, num2 int) int {

	return num1 + num2
}

func proAdder(nums ...int) (int, string) {

	total := 0

	for _, value := range nums {
		total += value
	}

	return total, "Hi from ProAdder"
}
