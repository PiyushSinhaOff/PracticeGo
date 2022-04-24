package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")

	var ptr *int
	fmt.Println(ptr)

	myNumber := 23
	fmt.Println(&myNumber)

	ptr = &myNumber
	fmt.Println(ptr, *ptr)

	*ptr = *ptr * 2
	fmt.Println(myNumber, *ptr)
}
