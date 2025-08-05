package main

import (
	"github.com/DustinMeyer1010/Learning/internal/tree"
)

func main() {
	array := []int{3, 19, 1, 14, 8, 7}
	root := tree.BuildBinaryTree(array)

	tree.PrintTree(root)

}
