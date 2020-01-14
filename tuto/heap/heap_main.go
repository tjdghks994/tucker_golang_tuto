package main

import (
	"fmt"
	heap "tucker_golang_tuto/tuto/datastruct"
)

func main() {
	var heap *heap.MinHeap = &heap.MinHeap{}

	// Q: [-1, 3, -1, 5, 4] 2번 째로 큰 수

	nums := []int{-1, 3, -1, 5, 4, 6}

	for i := 0; i < len(nums); i++ {
		heap.Push(nums[i])
		if heap.Count() > 2 {
			heap.Pop()
		}
	}
	fmt.Println(heap.Pop())

}
