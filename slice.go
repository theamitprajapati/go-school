package main

import (
	"fmt"
)

func main() {

	var arr []int
	arr = make([]int, 5)
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println("Typeof slice ", (arr))

	st := make([]int, 1)
	fmt.Println(st)

	copy(st, arr)
	fmt.Println("COPY:", arr, st)

	fmt.Println("range:", arr[:3])
	fmt.Println("from and to :", arr[2:3])
}
