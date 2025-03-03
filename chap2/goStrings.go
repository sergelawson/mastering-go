package main

import (
	"fmt"
	"strconv"
	s "strings"
	"unicode/utf8"
)

var pf = fmt.Printf
var pl = fmt.Println

func goStrings() {

	aString := "Hello y'all, I gotta get the bag 💰"

	//	sBytes := []byte(aString)

	pl(aString)

	for _, v := range aString {
		pf("%x ", v)
	}

	pl("")

	for _, v := range aString {
		pf("%c", v)
	}

	pl()

	d := string(128293)

	pf("%s\n", d)

	sNumber := strconv.Itoa(128293)

	pl(sNumber)

	pf("To Upper: %s\n", s.ToUpper("Hello There"))

	pf("To Lower %s\n", s.ToLower("UPPER CASE TEXT"))

	pf("Are equal %v\n", s.EqualFold("Lamborghini URUS", "LAMBORGHINI URUS"))

	pf("Index of UR %v\n", s.Index("Lamborghini URUS", "UR"))

	pf("Index of ur %v\n", s.Index("Lamborghini URUS", "ur"))

	pf("Repaced string: %s\n", s.Replace("heroin is not good", "heroin", "coke", -1))

	pf("number of word in string: %d\n", s.Count("Word of word", "word"))

	fields := s.Fields("Word for word")

	pf("%v\n", fields)

	strr := "relvo!"

	runeCount := utf8.RuneCountInString(strr)

	pf("number of characters %d\n", runeCount)

	pf("%v", s.Split("hell no, I'aint going to do that!", " "))

	/*
		str := string("\x99\x00ab\x50\x00\x23\x50\x29\x9c")

		r, _ := utf8.DecodeRuneInString(str)
		if r == utf8.RuneError {
			fmt.Printf("Invalid UTF-8 sequence at position: %v\n", len(str))

		}

		pf("cc %c %x", r)

		for _, r := range str {
			x := rune(r)
			if unicode.IsPrint(x) {
				pl("Rune is not printable", x)
				continue
			}
			pf("%U\n", x)
		}
	*/
	// pl(rList)
}
