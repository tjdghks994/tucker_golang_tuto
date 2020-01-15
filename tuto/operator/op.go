package main

import "fmt"

func main() {
	a := 4

	//fmt.Printf("%v\n", a&b)
	//fmt.Printf("%v\n", a|b)

	if (a & (a - 1)) == 0 {
		fmt.Println("a는 2의 거듭제곱")
	} else {
		fmt.Println("a는 거듭제곱이 아님")
	}

	var b bool = 3 < 4

	fmt.Println(b)
}
