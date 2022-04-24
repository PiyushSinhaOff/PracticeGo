package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case Example with some Random Numbers")

	rand.Seed(time.Now().UnixNano())
	dice := rand.Intn(6) + 1
	fmt.Println("Value of dice is :", dice)

	switch dice {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Dice value is 2 and you can not open")
	case 3:
		fmt.Println("Dice value is 3 and you can not open")
		fallthrough
	case 4:
		fmt.Println("Dice value is 4 and you cannot open")
	case 5:
		fmt.Println("Dice value is 5 and you cannot open")
	case 6:
		fmt.Println("Dice value is 6 and you can not open")
	}
}
