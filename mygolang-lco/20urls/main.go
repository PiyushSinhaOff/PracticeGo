package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://httpstat.us:4321/learn?coursename=Golang&paymentid=qwerty"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myUrl)

	// parsing the url
	result, err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println("result", result.Scheme)
	fmt.Println("result", result.Host)
	fmt.Println("result", result.Path)
	fmt.Println("result", result.Port())
	fmt.Println("result", result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type of qparams : %T \n", qparams)

	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	// better way

	for key, value := range qparams {
		fmt.Printf("Key is %v = %v \n", key, value)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/training",
		RawPath: "username=Piyush",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
