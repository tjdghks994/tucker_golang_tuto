package main

import (
	"fmt"
)

func main() {

	/*
		var a int = 30

		switch 30 {
		case a - 1:
			fmt.Println("a는 31")
		case a - 2:
			fmt.Println("a는 32")
		case a - 3:
			fmt.Println("a는 33")
		default:
			fmt.Println("a는 30")
		}
	*/

	var sum int = 0

	for j := 0; j < 10; j++ {
		if j == 8 {
			continue
		}
		fmt.Println(j, "입니다")
		sum += j

	}

	fmt.Println("1에서 10까지의 합은 ", sum, "입니다")

}
