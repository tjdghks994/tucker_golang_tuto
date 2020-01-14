package datastruct

import "fmt"

type BinaryTreeNode struct {
	left  *BinaryTreeNode //BST에서 왼쪽 노드는 자기보다 작은 값을 갖음
	right *BinaryTreeNode //BST에서 오른쪽 노드는 자기보다 큰 값을 갖음

	Val int
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func NewBinaryTree(v int) *BinaryTree {
	var tree *BinaryTree = &BinaryTree{}
	tree.Root = &BinaryTreeNode{Val: v}
	return tree
}

func (n *BinaryTreeNode) AddNode(v int) *BinaryTreeNode {
	if n.Val > v {
		if n.left == nil {
			n.left = &BinaryTreeNode{Val: v}
			return n.left
		} else {
			return n.left.AddNode(v)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryTreeNode{Val: v}
			return n.right
		} else {
			return n.right.AddNode(v)
		}
	}
}

type depthNode struct {
	depth int
	node  *BinaryTreeNode
}

func (t *BinaryTree) PrintTree() {
	q := []depthNode{}
	q = append(q, depthNode{depth: 0, node: t.Root})
	currentDepth := 0

	for len(q) > 0 {
		var first depthNode
		first, q = q[0], q[1:]

		if first.depth != currentDepth {
			fmt.Println()
			currentDepth = first.depth
		}
		fmt.Print(first.node.Val, " ")

		if first.node.left != nil {
			q = append(q, depthNode{depth: currentDepth + 1, node: first.node.left})
		}
		if first.node.right != nil {
			q = append(q, depthNode{depth: currentDepth + 1, node: first.node.right})
		}
	}
}

func (t *BinaryTree) SearchTree(v int) (bool, int) {
	return t.Root.SearchNode(v, 1)
}

func (n *BinaryTreeNode) SearchNode(v int, cnt int) (bool, int) {
	if n.Val == v {
		return true, cnt
	} else if n.Val > v {
		if n.left != nil {
			return n.left.SearchNode(v, cnt+1)
		}
		return false, cnt
	} else {
		if n.right != nil {
			return n.right.SearchNode(v, cnt+1)
		}
		return false, cnt
	}
}
