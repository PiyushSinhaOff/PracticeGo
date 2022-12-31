package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Something here please : ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Your something is here", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(numRating + 1)
}

// https://pkg.go.dev/strconv
