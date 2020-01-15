package main

import "fmt"

func main() {
	a := []int{100, 200, 300}

	fmt.Println(a)
	fmt.Println(len(a))               // 현재 갖고 있는 elements의 갯수
	fmt.Println(cap(a))               // 최대 넣을 수 있는 elements의 갯수
	fmt.Printf("%p, %p\n", &a, &a[0]) // Slice는 배열의 시작 주소를 포인터 값으로 갖는 구조체
	// 그러므로 slice의 주소와 배열의 시작 주소는 다름

	a = append(a, 400, 500, 600, 700)

	fmt.Println(a, cap(a))

	s := make([]int, 1)

	for i := 0; i < len(a); i++ {
		s = append(s, a[i])
	}

	fmt.Println(s)

	sh := student{"sh", 27, "A"}

	fmt.Println(sh)
}

type student struct {
	name  string
	age   int
	grade string
}
