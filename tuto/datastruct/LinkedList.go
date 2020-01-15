// Package datastruct 대문자: Public 소문자: prinvate
package datastruct

import "fmt"

// Node double Linked List node 구조체
type Node struct {
	Next *Node
	Prev *Node
	Val  int
}

// LinkedList double Linked List
type LinkedList struct {
	Root *Node
	Tail *Node
}

// AddNode add node
func (l *LinkedList) AddNode(val int) {
	if l.Root == nil {
		l.Root = &Node{Val: val}
		l.Tail = l.Root
		return
	}

	l.Tail.Next = &Node{Val: val}
	prev := l.Tail
	l.Tail = l.Tail.Next
	l.Tail.Prev = prev
}

// Back used by stack
func (l *LinkedList) Back() int {
	if l.Tail != nil {
		return l.Tail.Val
	}
	return 0
}

// PopBack used by stack
func (l *LinkedList) PopBack() {
	if l.Tail == nil {
		return
	}
	l.RemoveNode(l.Tail)
}

// Empty is Empty
func (l *LinkedList) Empty() bool {
	return l.Root == nil
}

// Front used by Queue
func (l *LinkedList) Front() int {
	if l.Root != nil {
		return l.Root.Val
	}
	return 0
}

// PopFront used by Queue
func (l *LinkedList) PopFront() {
	if l.Root == nil {
		return
	}
	l.RemoveNode(l.Root)
}

// RemoveNode remove node
func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.Root {
		l.Root = l.Root.Next
		if l.Root != nil {
			l.Root.Prev = nil // access violation
		}
		node.Next = nil
		return
	}

	prev := node.Prev

	if node == l.Tail {
		prev.Next = nil
		l.Tail.Prev = nil
		l.Tail = prev
	} else {
		node.Prev = nil
		prev.Next = prev.Next.Next
		prev.Next.Prev = prev
	}
	node.Next = nil
}

// PrintNode print node
func (l *LinkedList) PrintNode() {

	node := l.Root

	for node.Next != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Printf("%d \n", node.Val)
}

// PrintReverse print reverse
func (l *LinkedList) PrintReverse() {
	node := l.Tail

	for node.Prev != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Prev
	}
	fmt.Printf("%d \n", node.Val)
}
