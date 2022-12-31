package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	// no inheritance in golang; No super or parent

	info := User{"Piyush", "piyush@gmail.com", true, 16}
	fmt.Println(info)
	fmt.Printf("Info details are: %+v \n", info)
	fmt.Printf("Info details are: %v \n", info)
	fmt.Printf("Name is %v", info.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
