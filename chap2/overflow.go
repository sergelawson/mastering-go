package main

import (
	"fmt"
	"math"
)

func overflow() {
	var i int
	for {

		if i == math.MaxInt {
			fmt.Print("Max int :", i)
			break
		}
		i += 1

	}
}
