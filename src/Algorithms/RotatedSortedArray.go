package algorithms

/**
Problem : https://leetcode.com/problems/search-in-rotated-sorted-array/
Level : Medium
Result :
Runtime: 0 ms, faster than 100.00% of Go online submissions for Search in Rotated Sorted Array.
Memory Usage: 2.6 MB, less than 95.88% of Go online submissions for Search in Rotated Sorted Array.

195 / 195 test cases passed.
Status: Accepted
Runtime: 0 ms
Memory Usage: 2.6 MB
 */
//func main() {
//	//input := []int{4,5,6,7,0,1,2}
//	//input := []int{4,5,6,7,0,1,2}
//	input := []int{3, 1}
//	targetNum := 1;
//	result := findIndexOfNum(&input, &targetNum)
//	fmt.Println(result)
//}

func findIndexOfNum(input *[]int, target *int) int {
	startIndex := 0
	endIndex := len(*input) - 1
	return findTarget(input, startIndex, endIndex, target)
}

func findTarget(input *[]int, startIndex int, endIndex int, target *int) int {
	if startIndex == endIndex {
		//check if this current num matches the target
		if (*input)[startIndex] == *target {
			return startIndex
		} else {
			return -1
		}
	}else if startIndex < endIndex {
		mid := startIndex +  (endIndex - startIndex) / 2
		if (*input)[mid] == *target {
			return mid
		}
		//break the problem

		if (*input)[mid] < (*input)[endIndex] {
			//mid to right end is sorted
			if *target >= (*input)[mid] && *target <= (*input)[endIndex] {
				return findTarget(input, mid + 1, endIndex, target)
			}else {
				return findTarget(input, startIndex, mid - 1, target)
			}
		} else {
			if *target <= (*input)[mid] && *target >= (*input)[startIndex]{
				return findTarget(input, startIndex, mid - 1, target)
			}else {
				return findTarget(input, mid + 1, endIndex, target)
			}
		}
	}
	return -1
}


