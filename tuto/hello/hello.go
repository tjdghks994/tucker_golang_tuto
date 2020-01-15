package main

import "fmt"

func main() {

	hello()

	var a int
	var b int
	a = 3
	b = 4

	fmt.Println(a + b)
}

func hello() {

	fmt.Println("hello 월드!")
}
