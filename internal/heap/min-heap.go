package heap

import "fmt"

type MinHeap struct {
	data []int
}

func (h *MinHeap) BuildHeap(arr []int) {
	h.data = arr

	for i := len(h.data)/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}
}

func (h *MinHeap) heapify(i int) {
	n := len(h.data)
	smallest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && h.data[left] < h.data[smallest] {
		smallest = left
	}

	if right < n && h.data[right] < h.data[smallest] {
		smallest = right
	}

	if smallest != i {
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		h.heapify(smallest)
	}
}

func (h *MinHeap) Print() {
	fmt.Println(h.data)
}
