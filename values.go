package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Int:", 1)
	fmt.Println("Int:", (154444444444))
	fmt.Println("Float/Double", 1.2555)
	fmt.Println("LongDouble", 1555555555555555555.888888888888888888888888888888)
	fmt.Println("Boolean:", true)
	fmt.Println("String:", "1")
	fmt.Println("Boolean:", false)
	fmt.Println(reflect.TypeOf([]int{}))
}
