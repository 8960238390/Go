package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to web module")

	response, err := http.Get(url)

	checkNilError(err)

	fmt.Printf("Response is of type : %T\n", response)

	databytes, err := ioutil.ReadAll(response.Body)

	checkNilError(err)

	content := string(databytes)

	fmt.Println(content)

	defer response.Body.Close()
}

func checkNilError(err error) {

	if err != nil {
		panic(err)
	}
}
