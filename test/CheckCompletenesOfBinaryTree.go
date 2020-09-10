package main

import (
	algorithms "Algorithms"
	"fmt"
	. "models"
)

func main() {
	rootNode := TreeNode{Val: 1}
	rootNode.Left = &TreeNode{Val: 2}
	rootNode.Left.Left = &TreeNode{Val: 4}
	rootNode.Left.Right = &TreeNode{Val: 5}
	rootNode.Right = &TreeNode{Val: 3}

	//rootNode.Right.Left = &TreeNode{Val: 7}
	fmt.Println(algorithms.IsCompleteTree(&rootNode))
}

