package algorithms

/**
Leetcode Problem : Dutch Problem
https://leetcode.com/problems/sort-colors/
Level Medium
Result:
Runtime: 0 ms, faster than 100.00% of Go online submissions for Sort Colors.
Memory Usage: 2.1 MB, less than 85.23% of Go online submissions for Sort Colors.
87 / 87 test cases passed.
Status: Accepted
Runtime: 0 ms
Memory Usage: 2.1 MB
Submitted: 1 minute ago

 */
func SortNumbers(data *[]int) {
	lowIndex := 0
	middleIndex := 0
	highIndex := len(*data) - 1

	for ; middleIndex <= highIndex; {
		if (*data)[middleIndex] == 0 {
			//swap lowIndex and middle index
			//increment both
			temp := (*data)[middleIndex]
			(*data)[middleIndex] = (*data)[lowIndex]
			(*data)[lowIndex] = temp
			middleIndex++
			lowIndex++
		} else if (*data)[middleIndex] == 1 {
			middleIndex++
		} else if (*data)[middleIndex] == 2 {
			//swap middleIndex and highIndex
			//decrement highIndex
			temp := (*data)[middleIndex]
			(*data)[middleIndex] = (*data)[highIndex]
			(*data)[highIndex] = temp
			highIndex--
		}
	}
}
