package main

import "fmt"

func intSeq() func() int {

	var x int = 0

	return func() int {
		x++
		return x
	}

}

func main() {

	number := intSeq()

	fmt.Println(number())
	fmt.Println(number())
	fmt.Println(number())
	fmt.Println(number())
	fmt.Println(number())
	fmt.Println(number())
	fmt.Println(number())
	fmt.Println(number())
	fmt.Println(number())
	fmt.Println(number())
	fmt.Println(number())
	fmt.Println(number())
	fmt.Println(number())

	funcCall := intSeq()
	fmt.Println(funcCall())

}
