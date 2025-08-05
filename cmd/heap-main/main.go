package main

import (
	"github.com/DustinMeyer1010/Learning/internal/heap"
)

func main() {
	minHeap := heap.MinHeap{}

	minHeap.BuildHeap([]int{4, 19, 1, 14, 8, 7})

	minHeap.Print()
}
