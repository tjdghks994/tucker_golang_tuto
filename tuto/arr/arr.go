package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := [10]int{}
	temp := [10]int{}

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(10)
	}
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		idx := arr[i]
		temp[idx]++
	}

	idx := 0
	for i := 0; i < len(temp); i++ {
		for j := 0; j < temp[i]; j++ {
			arr[idx] = i
			idx++
		}
	}

	fmt.Println(arr)

}

//Radix sort: O(n)인 정렬 방법
//단, 배열 안의 값의 범위가 한정적이며 작을 때 사용 가능
//예, 이름의 앞 글자를 사용하여 정렬하는 경우

