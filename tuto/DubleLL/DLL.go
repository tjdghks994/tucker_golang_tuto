package main

import (
	"fmt"

	ds "github.com/tjdghks994/tucker_golang_tuto/tuto/datastruct"
)

func main() {
	var dll *ds.LinkedList = &ds.LinkedList{}
	dll.AddNode(0)

	for i := 1; i < 10; i++ {
		dll.AddNode(i)
	}

	dll.PrintNode()

	dll.RemoveNode(dll.Root.Next)

	dll.PrintNode()

	fmt.Println("now tail = ", dll.Tail.Val)

	dll.PrintReverse()

	/*
		array.append() = O(N) (단, cap이 여유로울 때는 O(1))
		array.remove() = O(N)
			마지막 element 가 삭제 될 때는 len의 pointer 만 하나 앞으로 당김 = O(1)
			중간에서 없앨 때 a.append([:n]) + a.append([n+1:]) 형식으로 이루어짐
			시작점에서 하나 지울 때는 index pointer 만 하나 뒤로 미룸 = O(1)
		배열은 추가, 삭제가 Linked List 보다 느리다
		단, 특정 인덱스에 접근하고 싶을 때(Random Access) 속도가 빠름 = O(1) // Linked List = O(N)
		Linked List => cache 를 추출 할 때 불리함
			cache 는 필요한 부분 주변의 연속된 부분을 가져옴, 그러나 Linked List는 불규칙하게 저장 되어있음
			매번 새롭게 cache를 가져와야함(cache miss 발생) 즉, 노드를 읽을 때 매번 새롭게 읽어들임

		정리
			*array는 데이터를 추가하거나 삭제할 때 느리다, 다만 특정 index를 읽어올 때는 빠르다
			*Linked List는 데이터를 추가하거나 삭제할 때는 빠르다, 다만 특정 index를 읽어올 때 느리다,
				또한 cache miss가 발생할 가능성을 갖고 있다.
	*/
	a := []int{1, 2, 3, 4, 5}
	a = append(a[:2], a[3:]...)
	fmt.Println(a)
}
