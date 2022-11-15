package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to our Pizza making app")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter number of pizza you want")
	number, err1 := reader.ReadString('\n')

	if err1 != nil {
		fmt.Println("err1 is : ", err1)
	} else {

		order, err2 := strconv.ParseInt(strings.TrimSpace(number), 10, 32)

		if err2 != nil {
			fmt.Println("err2 is : ", err2)
		} else {
			fmt.Println("value after conversion", order+1)
		}
	}
}
