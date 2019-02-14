package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		str := ""
		switch {
		case i%3 == 0 && i%5 == 0:
			str = "Fizz Buzz"
		case i%3 == 0:
			str = "Fizz"
		case i%5 == 0:
			str = "Buzz"
		default:
			str = fmt.Sprintf("%d", i)
		}
		fmt.Println(str)

	}
}
