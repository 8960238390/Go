package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paytemtid=kjsbd51sds"

func main() {

	fmt.Println("Welcome to hadling url")

	// accesing url diffn-diffn parts

	fmt.Println(myUrl)

	//parsing
	result, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Port())
	// fmt.Println(result.Path)
	// fmt.Println(result.RawQuery)

	qParam := result.Query()
	fmt.Printf("The type of parameter is %T \n", qParam)
	fmt.Println(qParam)

	for key, val := range qParam {
		fmt.Println(key, val)
	}

	// url construction

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	fmt.Println(partsOfUrl)
	fmt.Println(partsOfUrl.String())
}
