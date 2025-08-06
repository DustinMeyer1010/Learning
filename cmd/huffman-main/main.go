package main

import "github.com/DustinMeyer1010/Learning/internal/huffman"

func main() {
	original := "abccdacffcdeee"

	freqTable := map[string]int{}

	for _, letter := range original {
		if _, exist := freqTable[string(letter)]; exist {
			freqTable[string(letter)] += 1
			continue
		}
		freqTable[string(letter)] = 1
	}

	array := []huffman.FrequencyNode{}

	for k, v := range freqTable {
		array = append(array, huffman.FrequencyNode{Letter: &k, Frequency: v})
	}
	root := huffman.CreateBinaryCode(array)
	huffman.Traverse(&root, "")
}
