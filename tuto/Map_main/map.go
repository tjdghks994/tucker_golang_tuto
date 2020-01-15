package main

import "fmt"

type Key struct {
	Key int
}

type Value struct {
	v int
}

func main() {
	var m map[string]string // [key]value
	// var m1 map[int]string
	// var m2 map[Key]*Value // 이 처럼 모든 형태가 가능

	m = make(map[string]string) // map 은 선언 후 init 을 해줘야 사용 할 수 있음

	m["abc"] = "bbb"

	fmt.Println(m["abc"])

	m1 := make(map[int]string)
	m1[53] = "53번"

	fmt.Println(m1[53])

	m2 := make(map[int]int)

	m2[10] = 5312
	m2[3] = 32
	m2[4] = 532

	fmt.Println(m2[10])
	fmt.Println(m2[3]) // m2[3]의 값이 없기 때문에 int의 기본 값이 0 이 출력 됨
	// string의 기본값 = 공백, bool의 기본값 = false

	m2[5] = 0
	v, ok := m2[5]   //ok 는 map에 값이 할당 되었는지 할당 되지 않았는지를 판별해줌
	v1, ok2 := m2[6] // m2[6] 은 할당 되지 않았기 때문에 ok값이 false

	fmt.Println(v, ok)
	fmt.Println(v1, ok2)

	delete(m2, 5) // m2 map에서 key값이 5인 map을 해제

	v, ok2 = m2[5]
	fmt.Println(m2[2], ok2)

	for key, value := range m2 { // 정렬 되어 나오지 않음, 무작위로 나옴
		fmt.Println(key, "= ", value)
	}
}
