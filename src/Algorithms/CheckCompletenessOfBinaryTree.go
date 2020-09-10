package algorithms

import (
	. "models"
)
/**
Problem :https://leetcode.com/problems/check-completeness-of-a-binary-tree/
Result :
119 / 119 test cases passed.
Status: Accepted
Runtime: 0 ms
Memory Usage: 3.3 MB

Runtime: 0 ms, faster than 100.00% of Go online submissions for Check Completeness of a Binary Tree.
Memory Usage: 3.3 MB, less than 34.21% of Go online submissions for Check Completeness of a Binary Tree.
 */
func IsCompleteTree(root *TreeNode) bool {
	return isCompleteTree(root)
}

func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode {}
	queue = append(queue, root)
	for ; len(queue) > 0; {
		node := queue[0]
		queue = queue[1:len(queue)]
		if node.Val == -1 {
			for _, val := range queue {
				if val.Val != -1 {
					return false
				}
			}
			return true
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		} else {
			queue = append(queue, &TreeNode{Val: -1})
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		} else {
			queue = append(queue, &TreeNode{Val: -1})
		}
	}
	return true
}