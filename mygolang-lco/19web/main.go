package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const url = "https://httpstat.us/"

func main() {
	fmt.Println("http response")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body) // Caller Responsibility

	fmt.Printf("Response type is : %T \n", response)

	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(dataBytes)

	fmt.Printf(content)
}
