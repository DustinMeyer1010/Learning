package huffman

import (
	"container/heap"
	"fmt"
)

type FrequencyNode struct {
	Letter    *string
	Frequency int
	Right     *FrequencyNode
	Left      *FrequencyNode
}

type HeapFreq []FrequencyNode

func (h HeapFreq) Len() int           { return len(h) }
func (h HeapFreq) Less(i, j int) bool { return h[i].Frequency < h[j].Frequency }
func (h HeapFreq) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *HeapFreq) Push(x any) {
	*h = append(*h, x.(FrequencyNode))
}

// Smallest value is at the end of the array as the largest element was moved to front
func (h *HeapFreq) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func CreateBinaryCode(frequencies []FrequencyNode) FrequencyNode {

	h := &HeapFreq{}
	for _, f := range frequencies {
		heap.Push(h, f)
	}
	for h.Len() != 1 {
		right := heap.Pop(h).(FrequencyNode)
		left := heap.Pop(h).(FrequencyNode)

		heap.Push(h, FrequencyNode{Letter: nil, Frequency: right.Frequency + left.Frequency, Right: &right, Left: &left})
	}

	return heap.Pop(h).(FrequencyNode)
}

func Traverse(node *FrequencyNode, prefix string) {

	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil && node.Letter != nil {
		fmt.Printf("%s: %s\n", *node.Letter, prefix)
	}

	Traverse(node.Left, prefix+"0")
	Traverse(node.Right, prefix+"1")

}
