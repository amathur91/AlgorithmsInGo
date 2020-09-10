package algorithms

import (
	. "models"
)
/**
Leetcode Problem : https://leetcode.com/problems/binary-tree-inorder-traversal/
Result:
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
Memory Usage: 2.1 MB, less than 41.59% of Go online submissions for Binary Tree Inorder Traversal.

68 / 68 test cases passed.
Status: Accepted
Runtime: 0 ms
Memory Usage: 2.1 MB
 */
func InorderTraverse(root *TreeNode) [] int{
	return inorderTraversal(root)
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	processedNode := make(map[*TreeNode]bool)
	queue := []*TreeNode{}
	queue = append(queue, root)

	for ; len(queue) > 0; {
		//pull from queue
		//if the node has right child push it to the queue[0]
		//mark the node as processed and set the value in map as true
		//if the node has left child push it to the queue[0]
		node := queue[len(queue) - 1]
		queue = queue[0:len(queue) - 1]
		processed, contains := processedNode[node]
		if !contains  || !processed{
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			processedNode[node] = true
			queue = append(queue, node)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}
		if processed {
			result = append(result, node.Val)
		}
	}
	return result
}
