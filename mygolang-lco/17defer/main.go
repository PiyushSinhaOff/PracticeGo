package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")

	myDefer()
}

func myDefer() {
	for i := 1; i < 6; i++ {
		defer fmt.Println(i)
	}
}
