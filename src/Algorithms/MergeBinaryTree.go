package algorithms
import (
	. "models"
)
/**
Results :
problem : https://leetcode.com/problems/merge-two-binary-trees/
183 / 183 test cases passed.
Status: Accepted
Runtime: 28 ms
Memory Usage: 8.6 MB

Runtime: 28 ms, faster than 91.43% of Go online submissions for Merge Two Binary Trees.
Memory Usage: 8.6 MB, less than 73.33% of Go online submissions for Merge Two Binary Trees.
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	newRoot := TreeNode{Val: t1.Val + t2.Val}
	merge(t1, t2, &newRoot)
	return &newRoot
}

func merge(t1 *TreeNode, t2 *TreeNode, newRoot *TreeNode) {
	if t1.Left != nil && t2.Left != nil {
		newRoot.Left = &TreeNode{Val: t1.Left.Val + t2.Left.Val}
		merge(t1.Left, t2.Left, newRoot.Left)
	} else {
		if t1.Left == nil {
			newRoot.Left = t2.Left
		} else {
			newRoot.Left = t1.Left
		}

	}


	if t1.Right != nil && t2.Right != nil {
		newRoot.Right = &TreeNode{Val: t1.Right.Val + t2.Right.Val}
		merge(t1.Right, t2.Right, newRoot.Right)
	} else {
		if t1.Right == nil {
			newRoot.Right = t2.Right
		} else {
			newRoot.Right = t1.Right
		}
	}
}