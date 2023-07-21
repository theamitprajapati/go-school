package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now().Weekday())
	var i = 4
	for i < 10 {

		fmt.Print(i)
		i++
	}

}
