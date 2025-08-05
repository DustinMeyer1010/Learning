package tree

import "fmt"

type Node struct {
	Val   int
	Right *Node
	Left  *Node
}

func BuildBinaryTree(array []int) *Node {

	if len(array) == 0 {
		return nil
	}

	root := Node{array[0], nil, nil}

	for _, val := range array[1:] {
		root.Add(val)
	}

	return &root

}

func (n *Node) Add(val int) {
	if n.Val > val && n.Right == nil {
		n.Right = &Node{Val: val}
		return
	}
	if n.Val <= val && n.Left == nil {
		n.Left = &Node{Val: val}
		return
	}

	if n.Val > val {
		n.Right.Add(val)
	}

	if n.Val <= val {
		n.Left.Add(val)
	}
}

func PrintTree(n *Node) {
	if n == nil {
		return
	}

	fmt.Println(n.Val)
	PrintTree(n.Right)
	PrintTree(n.Left)

}
