package main

import (
	"fmt"
	"models"
)
import algorithms "Algorithms"


func main() {
	rootNode := models.TreeNode{Val: 3, Left: nil, Right: nil}
	rootNode.Left = &models.TreeNode{Val: 5}
	rootNode.Left.Left = &models.TreeNode{Val: 6}
	rootNode.Left.Right = &models.TreeNode{Val: 2}
	rootNode.Left.Right.Left = &models.TreeNode{Val: 7}
	rootNode.Left.Right.Right = &models.TreeNode{Val: 4}
	rootNode.Right = &models.TreeNode{Val: 1}
	rootNode.Right.Left = &models.TreeNode{Val: 0}
	rootNode.Right.Right = &models.TreeNode{Val: 8}
	//rootNode := models.TreeNode{Val: 1}
	//rootNode.Left = &models.TreeNode{Val: 2}
	result := algorithms.LowestCommonAncestor(&rootNode, &rootNode, rootNode.Left)
	fmt.Println(result)

}
