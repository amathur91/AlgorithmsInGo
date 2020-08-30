package algorithms

/**
Problem : https://leetcode.com/problems/target-sum/
Level : Medium
Runtime: 772 ms, faster than 24.83% of Go online submissions for Target Sum.
Memory Usage: 2.1 MB, less than 93.29% of Go online submissions for Target Sum.
139 / 139 test cases passed.
Status: Accepted
Runtime: 772 ms
Memory Usage: 2.1 MB

 */
func findTargetSumWays(nums []int, S int) int {
	totalCount := 0
	calculateTotalCount(&nums, 0, 0,&S, &totalCount)
	return totalCount
}

func calculateTotalCount(nums *[]int, currentIndex int, totalSum int, targetSum *int, totalCount *int) {
	if currentIndex < len(*nums) {
		calculateTotalCount(nums, currentIndex + 1, totalSum + (*nums)[currentIndex], targetSum, totalCount)
		calculateTotalCount(nums, currentIndex + 1, totalSum + -1 * (*nums)[currentIndex], targetSum, totalCount)
	} else {
		if totalSum == *targetSum {
			*totalCount++
		}
	}
}