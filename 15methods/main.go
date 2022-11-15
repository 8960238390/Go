package main

import "fmt"

func main() {

	fmt.Println("Welcome to Stucts Example")

	sharif := User{"Sharif", "sharif@gmail.com", true, 28}
	fmt.Println(sharif)
	fmt.Printf("User details %+v \n", sharif)
	fmt.Printf("Name is %v and email is %v \n", sharif.Name, sharif.Email)

	sharif.GetStatus()
	sharif.NewMail()
	fmt.Printf("User details %+v \n", sharif)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {

	fmt.Println("Is user active : ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@tcs.com"
	fmt.Println("updated value for mail id is : ", u.Email)
}
