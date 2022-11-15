package main

import "fmt"

func main() {

	var name string = "Sharif"
	fmt.Println(name)
	fmt.Printf("variable type is %T \n", name)

	var isbooked bool = false
	fmt.Println(isbooked)
	fmt.Printf("variables type is %T \n", isbooked)

	var age uint8 = 255
	fmt.Println(age)
	fmt.Printf("variables type is %T \n", age)

	var smallFloat float32 = 13.2211554466887733
	fmt.Println(smallFloat)
	fmt.Printf("variables type is %T \n", smallFloat)

	var largeFloat float64 = 13.2211554466887733
	fmt.Println(largeFloat)
	fmt.Printf("variables type is %T \n", largeFloat)

	// default values and some aliases
	var anotherVariable float32
	fmt.Println(anotherVariable)
	fmt.Printf("variables type is %T \n", anotherVariable)

	var website = "youtube.com"
	fmt.Println(website)
	fmt.Printf("variables type is %T \n", website)

	//vulnarous opertaor (no need to of var keyword and type of operator to declare function)
	// this operator is only allowed in side the function
	website1 := "google.com"
	fmt.Println(website1)
	fmt.Printf("variables type is %T \n", website1)

}
