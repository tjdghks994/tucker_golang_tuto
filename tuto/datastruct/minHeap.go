package datastruct

import "fmt"

type MinHeap struct {
	list []int
}

func (h *MinHeap) Push(v int) {
	h.list = append(h.list, v)

	idx := len(h.list) - 1

	for idx >= 0 {
		parentIdx := (idx - 1) / 2
		if parentIdx < 0 {
			break
		}
		if h.list[idx] < h.list[parentIdx] {
			h.list[idx], h.list[parentIdx] = h.list[parentIdx], h.list[idx]
			idx = parentIdx
		} else {
			break
		}
	}
}

func (h *MinHeap) Pop() int {
	if len(h.list) == 0 {
		return 0
	}

	top := h.list[0]
	last := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]

	if len(h.list) == 0 {
		return top
	}

	h.list[0] = last
	idx := 0
	for idx < len(h.list) {
		swapIdx := -1

		leftIdx := idx*2 + 1
		if leftIdx >= len(h.list) { // 자식 노드가 없을 경우, 왼쪽이 없으면 당연히 오른쪽도 없다
			break
		}
		if h.list[leftIdx] < h.list[idx] { // 왼쪽 노드 값이 더 크면 왼쪽 노드의 값으로 바꿀 준비, 바로 안바꾸는 이유는 오른쪽 노드의 값과도 비교해봐야함
			swapIdx = leftIdx
		}

		rightIdx := idx*2 + 2
		if rightIdx < len(h.list) { //오른쪽 자식 노드가 있으면
			if h.list[rightIdx] < h.list[idx] { // 오른쪽 노드의 값이 현재 인덱스의 값보다 크다면
				if swapIdx < 0 || h.list[swapIdx] > h.list[rightIdx] { // 왼쪽 노드가 인덱스 노드보다 커서 swqpIdx 값이 생긴 경우, 오른쪽 노드가 왼쪽보다 크면 swqpIdx를 right로 바꾼다
					swapIdx = rightIdx
				}
			}
		}

		if swapIdx < 0 {
			break
		}
		h.list[idx], h.list[swapIdx] = h.list[swapIdx], h.list[idx]

		idx = swapIdx
	}

	return top
}

func (h *MinHeap) Count() int {
	return len(h.list)
}

func (h *MinHeap) PrintHeap() {
	fmt.Println(h.list)
}
