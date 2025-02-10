package main

import (
	"fmt"
	"os"
	"strconv"
)

var Global int = 1234
var AnotherGlobal = -5678

func learning_switch() {

	argument := os.Args[1]

	age, err := strconv.Atoi(argument)

	if err != nil {
		fmt.Println("Please enter a valid age!")
		return
	}

	switch {
	case age >= 18:
		fmt.Printf("You are %d years old, and you are eligible to pursue  your application\n", age)

	default:
		fmt.Println("You are under age, you need to join your parents concent document to your application")
	}

	// fmt.Printf("Global %d \n", Global)
}
