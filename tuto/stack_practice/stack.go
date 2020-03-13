package main

import (
	"fmt"

	ds "github.com/tjdghks994/tucker_golang_tuto/tuto/datastruct"
)

func main() {
	// Slice를 사용한 stack, queue
	Stack := []int{}

	for i := 0; i < 5; i++ {
		Stack = append(Stack, i)
	}

	fmt.Println("스택: ", Stack)

	for len(Stack) > 0 {
		var last int
		last, Stack = Stack[len(Stack)-1], Stack[:len(Stack)-1]
		fmt.Println(last, Stack)
	}

	fmt.Println("---------------------------------")

	Queue := []int{}

	for i := 0; i < 5; i++ {
		Queue = append(Queue, i)
	}

	fmt.Println("큐: ", Queue)

	for len(Queue) > 0 {
		var first int
		first, Queue = Queue[0], Queue[1:]
		fmt.Println(first, Queue)
	}

	Stack2 := ds.NewStack()
	fmt.Println("New Stack")

	for i := 0; i < 5; i++ {
		Stack2.Push(i)
	}

	for !Stack2.Empty() {
		val := Stack2.Pop()
		fmt.Printf("%d -> ", val)
	}
	fmt.Println("\n----------------------------------")

	Queue2 := ds.NewQueue()
	fmt.Println("New Queue")

	for i := 0; i < 5; i++ {
		Queue2.Push(i)
	}

	for !Queue2.Empty() {
		val := Queue2.Pop()
		fmt.Printf("%d -> ", val)
	}
}
