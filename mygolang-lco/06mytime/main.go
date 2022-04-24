package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Study Package")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// createdTime := time.Date(2023, time.January, 10, 16, 10, 0, 0, time.UTC)
	// fmt.Println(createdTime)
}
