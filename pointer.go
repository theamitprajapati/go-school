package main

import "fmt"

func zeroVal(val int) {
	val = 0
}

func zeroPointer(p *int) {
	*p = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroVal(i)
	fmt.Println("zeroval:", i)

	zeroPointer(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

}
