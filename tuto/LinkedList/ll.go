package main

import "fmt"

func main() {
	var root *Node
	var tail *Node

	root = &Node{val: 0}
	tail = root

	for i := 1; i < 10; i++ {
		tail = AddNode(tail, i)

	}

	PrintNode(root)

	//root, _ = RemoveNode(root, root, tail)

	//PrintNode(root)

	//root, _ = RemoveNode(tail, root, tail)

	//PrintNode(root)

	root, _ = RemoveNodeWithVal(5, root, tail)

	PrintNode(root)
}

type Node struct {
	next *Node
	val  int
}

// AddNode O(1), for 문을 돌지 않음
func AddNode(tail *Node, value int) *Node {
	node := &Node{val: value}
	tail.next = node
	return node
}

// RemoveNode O(N)
func RemoveNode(node *Node, root *Node, tail *Node) (*Node, *Node) {
	if node == root {
		root = root.next
		if root == nil {
			tail = nil
		}
		return root, tail
	}

	prev := root
	for prev.next != node {
		prev = prev.next
	}

	if node == tail {
		prev.next = nil
		tail = prev
	} else {
		prev.next = prev.next.next
	}

	return root, tail
}

func RemoveNodeWithVal(val int, root *Node, tail *Node) (*Node, *Node) {
	if val == root.val {
		root = root.next
		if root == nil {
			tail = nil
		}
		return root, tail
	}

	prev := root
	for prev.val != (val - 1) {
		prev = prev.next
	}

	if val == tail.val {
		prev.next = nil
		tail = prev
	} else {
		prev.next = prev.next.next
	}

	return root, tail
}

// PrintNode node 출력
func PrintNode(root *Node) {
	node := root
	for node.next != nil {
		fmt.Printf("%d ->", node.val)
		node = node.next
	}
	fmt.Println(node.val)
}
