package main

import "fmt"

func main() {
	var a int
	var b int
	var p *int

	p = &a
	a = 3
	b = 2

	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(*p)

	p = &b

	fmt.Println(p)
	fmt.Println(*p)

	x := 1
	Increase(&x)

	fmt.Println(x)
}

// Increase 포인터
func Increase(x *int) {
	*x = *x + 1
}
