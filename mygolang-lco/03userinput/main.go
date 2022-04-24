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
	if(err != nil){
		println("Error reading")
	}
	fmt.Println("Thanks for rating", input)
	fmt.Printf("Type of the rating is %T", input)
}
