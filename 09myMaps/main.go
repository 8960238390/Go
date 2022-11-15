package main

import "fmt"

func main() {

	fmt.Println("Welcome to Maps Example")

	languages := make(map[string]string)

	fmt.Println(languages)
	fmt.Printf("type of the map is %T \n", languages)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Pyhton"

	fmt.Println(languages)

	fmt.Println("JS Shorts for : ", languages["JS"])

	delete(languages, "RB")

	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("For keys %v, value is %v\n", key, value)
	}

	// if we not intersed in keys

	for _, value := range languages {
		fmt.Printf("For keys v, value is %v\n", value)
	}

	// similar way if we are intested in values
	for key, _ := range languages {
		fmt.Printf("For keys %v, value is v\n", key)
	}
}
