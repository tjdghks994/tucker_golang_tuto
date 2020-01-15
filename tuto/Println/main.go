package main

import (
	"fmt"
	"strconv"
)

type StructA struct {
	val string
}

func (a *StructA) String() string { // 해당 함수가 있으면 print 시 해당 함수의 return 값을 출력
	return "Val: " + a.val
}

type StructB struct {
	val int
}

func (b *StructB) String() string {
	b.val += 10 // 뭘 하든 상관 없음
	return "Struct B: " + strconv.Itoa(b.val)
}

// Printable method로 String() string 을 갖는 struct는 Printable interface로 묶이게 됨
type Printable interface {
	String() string
}

// Println Printable interface 로 관계를 맺고 있는 구조체면 모두 사용 가능
func Println(p Printable) { // String() 함수를 갖는 객체를 묶음, duck type coding 방식이므로 가능
	fmt.Println(p.String())
}

func main() {
	a := &StructA{val: "AAA"}
	fmt.Println(a) // &{AAA}

	Println(a)

	b := &StructB{val: 6}

	Println(b)
}
