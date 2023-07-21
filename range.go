package main

import (
	"fmt"
	"reflect"
)

func main() {

	var arr = [10]int{1, 2, 3, 4} //array
	var maps = map[string]int{"a": 1, "b": 2, "c": 3}
	var slices = []int{1, 3, 2}
	fmt.Println("====================================", "\n")
	for index, value := range arr {
		fmt.Println(reflect.ValueOf(arr).Kind(), "== ARRay:", "index:", index, "value:", value)
	}
	fmt.Println("====================================", "\n")
	for index, value := range maps {
		fmt.Println(reflect.ValueOf(arr).Kind(), " == Map:", "index:", index, "value:", value)
	}
	fmt.Println("====================================", "\n")
	for index, value := range slices {
		fmt.Println(reflect.ValueOf(slices).Kind(), " == SLICE:=", "index:", index, "value:", value)
	}

	for index, value := range "amitz!" {
		fmt.Println(index, value)
	}

}
