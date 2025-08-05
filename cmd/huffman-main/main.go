package main

import "github.com/DustinMeyer1010/Learning/internal/huffman"

func main() {
	a := "a"
	b := "b"
	c := "c"
	root := huffman.CreateBinaryCode([]huffman.FrequencyNode{{Letter: &a, Frequency: 1}, {Letter: &b, Frequency: 5}, {Letter: &c, Frequency: 3}})
	huffman.Traverse(&root, "")
}
