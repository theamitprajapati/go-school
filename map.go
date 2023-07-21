package main

import "fmt"

func main() {

	x := map[string]int{"age": 20, "name": 123, "": 0.00}
	a := map[string]string{}
	x["ages"] = 10

	fmt.Println(len(x), a, x["sd"] == 0)

	x["sd"] = 1000
	value, err := x["sd"]
	fmt.Println("err:", err)
	fmt.Println("x5:", value)
	fmt.Println(x)

}
