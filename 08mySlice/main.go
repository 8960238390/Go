package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to Slice introduction")

	// way 1 : of declaring Slices
	var fruits = []string{"Apple", "Tomoto", "Peach"}
	fmt.Printf("Type of fruits is %T \n", fruits)

	fmt.Println(fruits)

	// fruits[3] = "Mangos"  // -> this is not allowed we can't increase the size of slice from this method we can add new element though the append() method

	fruits = append(fruits, "Mangos", "Banana")

	fmt.Println(fruits)

	fmt.Println("append new ", append(fruits[:]))
	fmt.Println("append new ", append(fruits[1:]))
	fmt.Println("append new ", append(fruits[:2]))

	fmt.Println()

	// way 2 : of decalring slices
	highScores := make([]int, 4)

	highScores[0] = 222
	highScores[1] = 124
	highScores[2] = 548
	highScores[3] = 36
	//highScores[4] = 51  // -> not allowed

	fmt.Println(highScores)

	highScores = append(highScores, 111, 451)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)

	fmt.Println("Sorted slices : ", highScores)

	fmt.Println("Removing element from slices")

	var courses = []string{"java", "hibernate", "sql", "spring", "microservices", "restfull"}

	fmt.Println(courses)

	index := 2
	fmt.Println(append(courses[:index], courses[index+1:]...))
}
