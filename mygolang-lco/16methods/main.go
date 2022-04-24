package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	// no inheritance in golang; No super or parent

	info := User{"Piyush", "piyush@gmail.com", true, 16}
	fmt.Printf("%+v \n", info)
	info.GetEmail()
	fmt.Printf("%+v ", info)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Printf("Is User Active : %v\n", u.Status)
}

func (u User) GetEmail() {				// call for Variable
	u.Email = "YOLO@gmail.com"
	fmt.Printf("Is User Active : %v\n", u.Email)
}
