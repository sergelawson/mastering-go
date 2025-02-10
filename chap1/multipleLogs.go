package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func multipleLogs(){
	
	flag := os.O_APPEND | os.O_CREATE | os.O_WRONLY

	file, err := os.OpenFile("myLog.log", flag, 0644)

	if (err != nil){
		fmt.Println(err)
		os.Exit(0)
	}

	defer file.Close()

	writer := io.MultiWriter(file, os.Stderr)

	logger := log.New(writer, "myApp: ", log.LstdFlags | log.Lshortfile)

	logger.Println("Logoging to multiple files")
}