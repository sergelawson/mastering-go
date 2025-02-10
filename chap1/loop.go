package main

import (
	"fmt"
	"strings"
)

func leaning_loop() {

	var colors [3]string = [3]string{"RED", "GREEEN", "BLUE"}

	for _, color := range colors {

		fmt.Println(strings.ToLower(color))
	}

	i := 0

	for i <= 10 {

		fmt.Printf("Counting to %d \n", i)
		i++
	}

	for j := 10; j <= 20; j++ {

		fmt.Printf("Counting to %d \n", j)

	}
}
