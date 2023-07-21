package main

import "fmt"

var total int = 1

func fact(num int) int {

	if num == 0 {
		return total
	}

	total *= num
	return fact(num - 1)
}

func main() {

	tt := fact(5)

	fmt.Println("Total fact:", tt)

}
