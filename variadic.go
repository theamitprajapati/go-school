package main

import (
	"fmt"
	"reflect"
	"strings"
)

func alphabet(str ...string) string {

	fmt.Println("Type:", reflect.ValueOf(str).Kind())
	fmt.Println(str)
	return strings.Join(str, ",")
}

func main() {

	var slices = []string{"am", "bb", "cb", "d", "e"}
	x := alphabet(slices...)

	fmt.Println(x)

	fmt.Println("hellowodk")
}
