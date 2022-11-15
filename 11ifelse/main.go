package main

import "fmt"

func main() {

	fmt.Println("if else Example")

	maritalStatus := false

	if maritalStatus {
		fmt.Println("you are married")
	} else {
		fmt.Println("you are not married")
	}

	age := 28

	if age < 10 {
		fmt.Println("you are a child")
	} else if age >= 10 && age < 20 {
		fmt.Println("you are a teen")
	} else {
		fmt.Println("you are an adult")
	}

	if 9%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10", num)
	} else {
		fmt.Println("Number is greater than 10", num)
	}
}
