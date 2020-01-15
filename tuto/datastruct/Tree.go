package datastruct

import "fmt"

// TreeNode Tree에서 사용 되는 Node
type TreeNode struct {
	Val      int
	Children []*TreeNode
}

// Tree Folder(Directory)를 구성할 때 주로 사용됨
type Tree struct {
	Root *TreeNode
}

//AddNode add
func (t *Tree) AddNode(val int) {
	if t.Root == nil {
		t.Root = &TreeNode{Val: val}
	} else {
		t.Root.Children = append(t.Root.Children, &TreeNode{Val: val})
	}
}

// AddNode children add
func (t *TreeNode) AddNode(val int) {
	t.Children = append(t.Children, &TreeNode{Val: val})
}

// DFSRecursive dfs call
func (t *Tree) DFSRecursive() {
	DFSRecursive(t.Root)
}

// DFSRecursive 재귀를 이용한 DFS, 길찾기 알고리즘 Dijkstra Algorithm, A* Algorithm
func DFSRecursive(t *TreeNode) {
	fmt.Printf("%d -> ", t.Val)

	for i := 0; i < len(t.Children); i++ {
		DFSRecursive(t.Children[i])
		//fmt.Printf("%d -> ", t.Children[i].Val) // 밑에서 부터 올라오기
	}
}

// DFSStack stack을 이용한 DFS
func (t *Tree) DFSStack() {
	stack := []*TreeNode{}
	stack = append(stack, t.Root)

	for len(stack) > 0 {
		var last *TreeNode
		last, stack = stack[len(stack)-1], stack[:len(stack)-1]

		fmt.Printf("%d -> ", last.Val)
		// 이 부분에 필요한 연산을 넣으면 됨

		for i := len(last.Children) - 1; i >= 0; i-- {
			stack = append(stack, last.Children[i])
		}
		/*
			for i := 0; i < len(last.Children); i++ {
				stack = append(stack, last.Children[i])
			}
		*/
	}
}

// BFSQueue Queue를 이용한 BFS
func (t *Tree) BFSQueue() {
	queue := []*TreeNode{}
	queue = append(queue, t.Root)

	for len(queue) > 0 {
		var first *TreeNode
		first, queue = queue[0], queue[1:]

		fmt.Printf("%d -> ", first.Val)
		// 이 부분에 필요한 연산을 넣으면 됨

		for i := 0; i < len(first.Children); i++ {
			queue = append(queue, first.Children[i])
		}

	}
}
