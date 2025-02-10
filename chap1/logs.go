package main

import (
	"log"
	"os"
)

func learning_logs() {

	args := os.Args

	if len(args) == 1 {
		log.Fatal("Fatal: arguments required.")
	}

	log.Panic("Panic: Hello World")

}
