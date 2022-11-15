package main

import "fmt"

func main() {

	fmt.Println("Welcome to loops")
	//days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// for i := 0; i < len(days); i++ {

	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for input, day := range days {
	// 	fmt.Printf("Element Index is %v and Element value is %v \n", input, day)
	// }

	input := 1
	for input < 10 {

		if input == 2 {
			goto lco
		}
		if input == 5 {
			input++
			continue
		}
		fmt.Printf("Input value is %v \n", input)
		input++
	}

lco:
	fmt.Println("Jumping at Learning codeonline.in")
}
