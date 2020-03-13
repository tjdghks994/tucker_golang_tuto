package main

import (
	"fmt"

	tr "github.com/tjdghks994/tucker_golang_tuto/tuto/datastruct"
)

func main() {
	var tree *tr.Tree = &tr.Tree{}

	val := 1
	tree.AddNode(val)
	val++

	for i := 0; i < 3; i++ {
		tree.Root.AddNode(val)
		val++
	}

	for i := 0; i < len(tree.Root.Children); i++ {
		for j := 0; j < 2; j++ {
			tree.Root.Children[i].AddNode(val)
			val++
		}
	}

	tree.DFSRecursive()
	fmt.Println()

	tree.DFSStack()
	fmt.Println()

	tree.BFSQueue()
	fmt.Println()
}
