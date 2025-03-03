package main

import (
	"fmt"
	"time"
)

func goDates() {

	now := time.Now().Local()

	parsedt, err := time.Parse("Jan 02 2006", "Feb 29 2024")

	if err != nil {

		fmt.Println("Invalid date")

		return
	}
	fmt.Println(parsedt, now)
	fmt.Println(parsedt.Format("le 02, January 2006"))

	location, err := time.LoadLocation("America/Los_Angeles")

	if err != nil {
		fmt.Println("Invalid time zone")
		return
	}
	fmt.Println(now.In(location).Format("15:04 Monday 02/01/2006 MST"))
}
