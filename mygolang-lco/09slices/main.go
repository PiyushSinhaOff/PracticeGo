package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var fruitList = []string{"Apple", "Orange", "banana"}
	fmt.Printf("Type of fruitList %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "beans")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:], "cake")
	fmt.Println(fruitList)

	highScoreList := make([]int, 4)

	highScoreList[0] = 100
	highScoreList[1] = 900
	highScoreList[2] = 300
	highScoreList[3] = 400

	fmt.Println(highScoreList)
	highScoreList = append(highScoreList, 500, 600)
	fmt.Println(highScoreList)

	fmt.Println(sort.IntsAreSorted(highScoreList))
	sort.Ints(highScoreList)
	fmt.Println(highScoreList)

	//Booleans values
	fmt.Println(sort.IntsAreSorted(highScoreList))


	//How to remove the value from slics on index 
	index := 1 
	highScoreList = append(highScoreList[:index], highScoreList[index+1:]...)
	fmt.Println(highScoreList)

	// How to remove the values from slics on values 
	
}
