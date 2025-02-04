package main

import "fmt"


func standard_input() {
	fmt.Println("Enter your name")

	var name string

	fmt.Scanln(&name)

	fmt.Printf("Welcome %s \n", name)
}