package main

import "fmt"

func main() {
	// var values = []string{"yolo", "polo", "tolo", "dolo"}

	// for i := 0; i < len(values); i++ {
	// 	fmt.Println(values[i])
	// }

	// for i := range values {
	// 	fmt.Println(values[i])
	// }

	// for ind, val := range values {
	// 	fmt.Println(ind, val)
	// }

	rougueInd := 1

	for rougueInd < 10 {
		// if rougueInd == 2 {
		// 	goto lco
		// }

		if rougueInd == 5 {
			rougueInd++
			continue
		}

		fmt.Println(rougueInd)
		rougueInd++
	}

// lco:
// 	fmt.Println("Jumping at Learncode")

}
