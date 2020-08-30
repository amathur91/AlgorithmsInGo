package algorithms

import "models"
/**
https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
Runtime: 376 ms, faster than 6.38% of Go online submissions for Lowest Common Ancestor of a Binary Tree.
Memory Usage: 7 MB, less than 90.14% of Go online submissions for Lowest Common Ancestor of a Binary Tree.

 */
func LowestCommonAncestor(root, p, q *models.TreeNode) *models.TreeNode {
	return lowestCommonAncestor(root, p, q)
}
func lowestCommonAncestor(root, p, q *models.TreeNode) *models.TreeNode {
	if root != nil {
		if root.Val == p.Val  || root.Val == q.Val{
			return root
		} else {
			result1 := findValueInTree(root.Left, &p.Val)
			result2 := findValueInTree(root.Left, &q.Val)
			if (result1 && !result2) || (!result1 && result2) {
				return root
			}
			if result1 && result2 {
				return lowestCommonAncestor(root.Left, p, q)
			}else {
				return lowestCommonAncestor(root.Right, p, q)
			}
		}
	}
	return nil
}

func findValueInTree(root *models.TreeNode, num *int) bool {
	if root != nil {
		if root.Val == *num {
			return true
		}
		return findValueInTree(root.Left, num) || findValueInTree(root.Right, num)
	}
	return false
}