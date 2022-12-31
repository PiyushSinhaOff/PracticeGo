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

	// Time in IST

	// First, we create an instance of a timezone location object
	loc, _ := time.LoadLocation("Asia/Kolkata")

	// this is our custom format. Note that the format must point to this exact time
	format := "Jan _2 2006 3:04:05 PM"

	// this is your timestamp
	timestamp := "Jun 25 2015 10:00:00 AM"

	// now we parse it, considering it's in IST
	t, err := time.ParseInLocation(format, timestamp, loc)

	// printing it prints it in IST, but you can set the timezone to UTC if you want
	fmt.Println(t, err)

	// example - getting the UTC timestamp
	fmt.Println(t.UTC())

	fmt.Println(time.Now().In(loc)) // CURRENT TIME IN IST

}
