package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func goRunes() {

	char := os.Args[1]

	//aString := "Hello y'all, I gotta get the bag 💰"
	//sBytes := []byte(aString)

	s := []rune(char)[0]

	// Create a buffer to hold the UTF-8 encoded bytes
	buf := make([]byte, utf8.RuneLen(s))

	// Encode the rune into the buffer
	utf8.EncodeRune(buf, s)

	// fmt.Println(s)

	fmt.Printf("character %c \nunicode %U \ntype %T\nhex %x \nbin %b\ndecimal %d\nbytes %v \n", s, s, s, s, s, s, buf)
	// fmt.Println(sBytes)
}
