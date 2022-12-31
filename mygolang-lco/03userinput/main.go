package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to USA"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	// comma ok || err ok
	input, err := reader.ReadString('\n')
	if err != nil {
		println("Error reading")
	}
	fmt.Println("Thanks for rating", input)
	fmt.Printf("Type of the rating is %T", input)

	// Another Way to Take input from the User
	// Println function is used to
	// display output in the next line
	fmt.Println("Enter Your First Name: ")

	// var then variable name then variable type
	var first string

	// Taking input from user
	fmt.Scanln(&first)
	fmt.Println("Enter Second Last Name: ")
	var second string
	check, err := fmt.Scanln(&second)
	fmt.Println(check)
	if err != nil {
		fmt.Println("okay")
	}

	// Print function is used to
	// display output in the same line
	fmt.Print("Your Full Name is: ")

	// Addition of two string
	fmt.Print(first + " " + second)

}
