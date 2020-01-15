package datastruct

// Stack stack memory: 다중 for 문 안에서 보통 나중에 선언된 변수가 먼저 없어짐
type Stack struct {
	LL *LinkedList
}

// NewStack init func
func NewStack() *Stack {
	return &Stack{LL: &LinkedList{}}
}

// Push Add
func (s *Stack) Push(val int) {
	s.LL.AddNode(val)
}

// Pop return last val
func (s *Stack) Pop() int {
	back := s.LL.Back()
	s.LL.PopBack()
	return back
}

// Empty is Empty
func (s *Stack) Empty() bool {
	return s.LL.Empty()
}
