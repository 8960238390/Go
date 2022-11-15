package main

import "fmt"

func main() {

	fmt.Println("Welcome to Stucts Example")

	sharif := User{"Sharif", "sharif@gmail.com", true, 28}
	fmt.Println(sharif)
	fmt.Printf("User details %+v \n", sharif)
	fmt.Printf("Name is %v and email is %v \n", sharif.Name, sharif.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
