package main

import (
	algorithms "Algorithms"
	"fmt"
	. "models"
)


func main() {
	root := TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	fmt.Println(algorithms.InorderTraverse(&root))
}