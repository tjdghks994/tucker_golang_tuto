package main

import "fmt"

func main() {
	//slice()
	//slice2()

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := 0
	for i := 0; i < 5; i++ {
		a, res = RemoveBack(a)
		fmt.Println(res, cap(a))
	}
	fmt.Println(a)
}

// RemoveBack 뒷 자리를 하나씩 지우는 함수
func RemoveBack(a []int) ([]int, int) {
	b := a[:len(a)-1]

	return b, a[len(a)-1]
}

func slice2() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(a[4:8], cap(a[4:8])) // [start : end] 는 start ~ end-1 까지 cap은 자른지점 부터 끝까지
	fmt.Println(a[4:])               // [start : ] 는 start 부터 마지막까지
	fmt.Println(a[:4])               // [ : end] 는 처음부터 end -1 까지

	b := a[4:8]
	b[0] = 1
	fmt.Println(a, b) // b는 a의 일부를 가르키고 있기 때문에 b를 바꿔도 a가 같이 바뀐다
}

func slice() { // slice 이해하기
	var a []int
	fmt.Println("len = ", len(a))
	fmt.Println("cap = ", cap(a))

	b := []int{1, 2, 3, 4}
	fmt.Println("len = ", len(b))
	fmt.Println("cap = ", cap(b))

	a = append(a, 1) // append(arr, int) 배열의 마지막에 값 추가
	fmt.Println("len = ", len(a))
	fmt.Println("cap = ", cap(a))

	b = append(b, 1)
	fmt.Println(b)

	a = append(b, 3) // a라는 배열에 b 배열에 3을 추가한 값을 대입
	fmt.Println(a)
	fmt.Println(b)

	c := []int{1, 2}
	d := []int{}

	d = append(c, 3)
	fmt.Println(c)
	fmt.Println(d)

	e := make([]int, 2, 4) // cap이 충분하여 같은 메모리의 주소를 가르키고 있음
	f := append(e, 3)
	fmt.Printf("%p, %p\n", e, f)

	f[0] = 4 // 따라서 f의 값만 바꿨음에도 e의 값도 같이 바뀜
	fmt.Println(e, f)

	// 따라서 위의 문제를 해결하기 위해서는 값을 복사 한 후 복사한 배열에 값을 추가하는 것이 안전함
	aa := []int{1, 2, 3, 4}
	//bb := []int{} 그냥 이렇게 하면 배열의 length가 맞지 않아 문제가 생김
	bb := make([]int, len(aa))

	for i := 0; i < len(aa); i++ { // 값을 복사
		bb[i] = aa[i]
	}
	bb = append(bb, 5)
	bb[0] = 6

	fmt.Println(aa, bb)
}
