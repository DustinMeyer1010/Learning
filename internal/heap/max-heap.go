package heap

type MaxHeap struct {
	data []int
}

func (h *MaxHeap) BuildHeap(arr []int) {
	h.data = arr
	for i := len(h.data)/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}
}

func (h *MaxHeap) heapify(i int) {
	n := len(h.data)
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && h.data[left] > h.data[largest] {
		largest = left
	}

	if right < n && h.data[right] > h.data[largest] {
		largest = right
	}

	if largest != i {
		h.data[i], h.data[largest] = h.data[largest], h.data[i]
		h.heapify(largest)
	}
}
