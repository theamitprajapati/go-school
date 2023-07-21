package main

import "fmt"

func main() {

	fmt.Println("Reading array implementation")

	var x [4]int

	fmt.Println("Len:", len(x))
	x[2] = -55

	fmt.Println(x)

}
