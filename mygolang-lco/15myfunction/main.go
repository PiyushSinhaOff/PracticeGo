package main

import "fmt"

func main() {
	val, multi := adder(3, 4, 8)
	fmt.Println(val, multi)
}

func adder(values ...int) (int, int) {
	total := 0
	totalMult := 1

	for _, v := range values {
		total += v
		totalMult *= v
	}

	return total, totalMult
}
