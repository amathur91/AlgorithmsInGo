package algorithms

import (
	"fmt"
	. "models"
)

/**
Invert a Binary Tree
Leetcode : https://leetcode.com/problems/invert-binary-tree
Results :
Runtime: 0 ms, faster than 100.00% of Go online submissions for Invert Binary Tree.
Memory Usage: 2.3 MB, less than 8.73% of Go online submissions for Invert Binary Tree.
68 / 68 test cases passed.
Status: Accepted
Runtime: 0 ms
Memory Usage: 2.3 MB
Submitted: 0 minutes ago
 */
//func main() {
//	root := models.TreeNode{Val: 4}
//	root.Left = &models.TreeNode{Val: 2}
//	root.Left.Left = &models.TreeNode{Val: 1}
//	root.Left.Right = &models.TreeNode{Val: 3}
//	root.Right = &models.TreeNode{Val: 7}
//	root.Right.Left = &models.TreeNode{Val: 6}
//	root.Right.Right = &models.TreeNode{Val: 9}
//	printTreeInOrder(&root)
//	invertedTree := createInvertedTree(&root)
//	fmt.Println()
//	printTreeInOrder(invertedTree)
//}

func createInvertedTree(root *TreeNode) *TreeNode {
	if root != nil {
		duplicateNode := &TreeNode{Val: root.Val}
		duplicateNode.Left = createInvertedTree(root.Right)
		duplicateNode.Right = createInvertedTree(root.Left)
		return duplicateNode
	}
	return nil
}

func printTreeInOrder(root *TreeNode) {
	if root != nil {
		printTreeInOrder(root.Left)
		fmt.Printf("%d ",root.Val)

		printTreeInOrder(root.Right)
	}
}
