package main

import (
	"errors"
	"fmt"
	"os"
)

func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a %d and b %d. UserID: %d", a, b, os.Getuid())
	}
	return nil
}

func check(a, b int) error {

	if a == 0 && b == 0 {

		return errors.New("custom error message")
	}

	return nil
}

func err() {

	err := check(0, 0)

	ferr := formattedError(0, 0)

	if err != nil {
		fmt.Println(err)
	}

	if ferr != nil {
		fmt.Println(ferr)
	}

}
