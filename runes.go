package main

import (
	"fmt"
	"unicode/utf8"
)

const s = "ดี"

func main() {

	var x = "Hello"

	for i := len(x) - 1; i > -1; i-- {
		a, b := utf8.DecodeRuneInString(x[i:])
		fmt.Printf("%c", a)
		fmt.Printf("%c", b)
	}
}
