package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func which() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Please enter aruments!")

		return
	}

	file := arguments[1]

	path := os.Getenv("PATH")

	pathList := filepath.SplitList(path)

	for _, directory := range pathList {

		fullpath := filepath.Join(directory, file)

		fileInfo, err := os.Stat(fullpath)

		if err != nil {
			continue
		}

		mode := fileInfo.Mode()

		if !mode.IsRegular() {
			continue
		}

		// Is it executable?
		if mode&0111 != 0 {
			fmt.Println(fullpath)
			return
		}
	}

}
