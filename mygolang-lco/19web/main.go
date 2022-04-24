package main

import (
	"fmt"
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

	defer response.Body.Close() // Caller Reponsibility

	fmt.Printf("Response type is : %T \n", response)

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)

	fmt.Printf(content)
}
