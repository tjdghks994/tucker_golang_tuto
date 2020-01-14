package main

import (
	"fmt"
	bst "tucker_golang_tuto/tuto/datastruct"
)

func main() {
	tree := bst.NewBinaryTree(5)
	tree.Root.AddNode(3)
	tree.Root.AddNode(2)
	tree.Root.AddNode(4)
	tree.Root.AddNode(8)
	tree.Root.AddNode(7)
	tree.Root.AddNode(6)
	tree.Root.AddNode(10)
	tree.Root.AddNode(9)

	tree.PrintTree()
	fmt.Println()

	n := 6

	//fmt.Scanln(&n)

	if found, cnt := tree.SearchTree(n); found {
		fmt.Println("found ", n, " cnt: ", cnt)
	} else {
		fmt.Println("not found ", n, " cnt: ", cnt)
	}
}
