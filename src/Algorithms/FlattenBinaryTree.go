package algorithms

import (
	"fmt"
	"models"
)

/**
Leetcode : https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
results :
Runtime: 8 ms, faster than 32.34% of Go online submissions for Flatten Binary Tree to Linked List.
Memory Usage: 7.9 MB, less than 5.39% of Go online submissions for Flatten Binary Tree to Linked List.

225 / 225 test cases passed.
Status: Accepted
Runtime: 8 ms
Memory Usage: 7.9 MB
*/
//func main() {
//	root := models.TreeNode{Val: 1}
//	root.Left = &models.TreeNode{Val: 2}
//	root.Left.Left = &models.TreeNode{Val: 3}
//	root.Left.Right = &models.TreeNode{Val: 4}
//	root.Right = &models.TreeNode{Val: 5}
//	root.Right.Right = &models.TreeNode{Val: 6}
//	linkedListRoot := flattenTree(&root)
//	printLinkedList(linkedListRoot)
//
//}

func flattenTree(root *models.TreeNode) *models.TreeNode {
	if root != nil {
		//flatten left subtree
		leftFlattenTree := flattenTree(root.Left)
		//flatten right subtree
		rightFlattenTree := flattenTree(root.Right)

		//find the right most of left flatten list
		tempNode := leftFlattenTree
		if tempNode != nil {
			for ; tempNode.Right != nil; tempNode = tempNode.Right {
			}
		}

		if root.Left != nil {
			//connect the right most of left to the head of right subtree list
			if tempNode != nil {
				tempNode.Right = rightFlattenTree
			}
			//final Adjustments
			root.Right = leftFlattenTree
		} else {
			root.Right = rightFlattenTree
		}
		root.Left = nil
		return root
	}
	return nil
}

func printLinkedList(node *models.TreeNode) {
	if node != nil {
		for tempNode := node; tempNode != nil; tempNode = tempNode.Right {
			fmt.Printf("%d ", tempNode.Val)
		}
	}
}
