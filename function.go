package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func errors() error {
	return fmt.Errorf("this is error")
}

func multipleReturn(x int) (int, string, error) {
	return x, "i am string", nil
}
func main() {

	x := sum(1, 2)

	fmt.Println(x)

	fmt.Println(errors())

	a, b, c := multipleReturn(3)

	fmt.Println(a, b, c)
}
