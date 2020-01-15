package main

import "fmt"

func main() {
	//dan := 3
	//gugu(dan)
	//fmt.Println(add(dan, dan))

	//a, b := change(2, 3)
	//fmt.Println(a, b)

	//rec(10)
	//sum := recsum(0, 10)
	//fmt.Println(sum)

	fibnum := fib(11)
	fmt.Println(fibnum)
}

func gugu(dan int) {
	for i := 1; i < 10; i++ {
		fmt.Println(dan, "x", i, "=", dan*i)
	}
}

func add(x int, y int) int {
	return x + y
}

func change(x, y int) (int, int) {
	return y, x
}

func rec(x int) {
	if x == 0 {
		return
	}
	fmt.Println(x)
	rec(x - 1)
	fmt.Println(x)
}

func recsum(x, y int) int {
	if y == 0 {
		return x
	}
	x += y
	fmt.Println(x)
	return recsum(x, y-1)
}

func fib(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	}

	return fib(x-1) + fib(x-2)
}
