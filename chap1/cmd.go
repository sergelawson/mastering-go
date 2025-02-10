package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

func calculate_density() {
	arguments := os.Args

	if len(arguments) < 3 {
		fmt.Println("Please enter the mass (-m) and the volume (-v) !")
		return
	}

	if len(arguments) == 3 {

		mass, errMass := strconv.ParseFloat(arguments[1], 64)
		volume, errVolum := strconv.ParseFloat(arguments[2], 64)

		if errMass != nil {

			fmt.Println("Valid mass required!")
			return
		}

		if errVolum != nil {
			fmt.Println("Valid volume required!")
			return
		}

		density := mass / volume

		fmt.Printf("The desity is %.2f \n", density)
		return
	}

	var massArg string
	var volumeArg string

	for i := 1; i < len(arguments); i++ {
		if utf8.RuneCountInString(massArg) > 0 && utf8.RuneCountInString(volumeArg) > 0 {
			break
		}
		if arguments[i] == "-m" {
			massArg = arguments[i+1]
		}
		if arguments[i] == "-v" {
			volumeArg = arguments[i+1]
		}
	}

	mass, errMass := strconv.ParseFloat(massArg, 64)
	volume, errVolum := strconv.ParseFloat(volumeArg, 64)

	if errMass != nil {

		fmt.Println("Valid mass required!")
		return
	}

	if errVolum != nil {
		fmt.Println("Valid volume required!")
		return
	}

	density := mass / volume

	fmt.Printf("The desity is %.2f \n", density)
}

// go run . 58 10
// go run . -m 20 -v 2
