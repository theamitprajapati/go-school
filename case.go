package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Data", time.Sunday)
	switch 3 {
	case 3:
		fmt.Println("time....")
	case 9:
		fmt.Println("Hell")
	default:
		fmt.Println("End switch case")
	}

	whoami := func(i interface{}) {
		switch i.(type) {

		case bool:
			fmt.Println("Boolean")
		case int:
			fmt.Println("Integer")
		case float64:
			fmt.Println("Float64")
		default:
			fmt.Println("Not implemented")
			//fmt.Println("data", i.(type))
		}
	}

	whoami("333")
	whoami(545454545454545)
	whoami(696)
	whoami(696.23)
	whoami(true)

}
