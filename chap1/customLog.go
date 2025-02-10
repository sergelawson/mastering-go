package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func customLog() {

	LOGFILE := path.Join(os.TempDir(), "customGo.log")
	fmt.Println(LOGFILE)

	file, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)

		return
	}

	defer file.Close()

	iLog := log.New(file, "iLog ", log.LstdFlags)

	iLog.Println("We are getting the logs here!")

	// Show line number
	iLog.SetFlags(log.Ldate | log.Lshortfile)

	iLog.Println("Show line number !")

	// Show file path and line number
	iLog.SetFlags(log.Ldate | log.Llongfile)

	iLog.Println("Show long line number !")

}
