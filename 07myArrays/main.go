package main

import "fmt"

func main() {

	fmt.Println("Welcome to Arrays concept")
	var fruits [4]string

	fmt.Println("values in variables : ", fruits)
	fmt.Println("Length : ", len(fruits))

	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[3] = "Pine Apple"
	fruits[3] = "Pine Apple"

	fmt.Println("values in variables : ", fruits)
	fmt.Println("Length : ", len(fruits))

	var vegitables = [5]string{"Beans", "Tomato", "Potato"}
	fmt.Println("values in variables : ", vegitables)
	fmt.Println("Length : ", len(vegitables))
	fmt.Printf("Variable type is %T \n", fruits)

}
