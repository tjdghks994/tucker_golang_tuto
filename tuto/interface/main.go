package main

import (
	"fmt"
	"strconv"
)

type StructA struct {
}

func (a *StructA) AAA(val int) int {
	return val * val
}

func (a *StructA) BBB(val int) string {
	return "X = " + strconv.Itoa(val)
}

type StructB struct {
}

func (b *StructB) BBB(val int) string {
	return "XB = " + strconv.Itoa(val)
}

func main() {
	var c InterFaceA
	c = &StructA{} // Struct A는 AAA(), BBB() 를 모두 갖고 있기 때문에 대입이 가능
	// c = &StructB{} // *StructB does not implement InterFaceA (missing AAA method) -> interface에 올라가있는 함수는 모두 구현이 되어있어야함!

	fmt.Println(c.BBB(3))
}
