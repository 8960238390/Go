package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to GO"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please provide your rating")

	// comma ok || err

	input, err := reader.ReadString('\n')
	//input, _ := reader.ReadString('\n')
	//_, err := reader.ReadString('\n')

	fmt.Println("Thanks for rating : ", input)
	fmt.Printf("Variable type is %T \n", input)
	fmt.Println("You are getting this error : ", err)
	fmt.Printf("Variable type is %T \n", err)

}
