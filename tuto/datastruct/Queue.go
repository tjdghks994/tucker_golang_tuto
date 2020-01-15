package datastruct

// Queue 대기열을 만들 때 주로 사용
type Queue struct {
	LL *LinkedList
}

// NewQueue init queue
func NewQueue() *Queue {
	return &Queue{&LinkedList{}}
}

// Push add
func (q *Queue) Push(val int) {
	q.LL.AddNode(val)
}

// Pop remove
func (q *Queue) Pop() int {
	val := q.LL.Front()
	q.LL.PopFront()
	return val
}

// Empty is Empty
func (q *Queue) Empty() bool {
	return q.LL.Empty()
}
