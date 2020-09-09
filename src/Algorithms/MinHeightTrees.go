package algorithms


/**

https://leetcode.com/problems/minimum-height-trees/
67 / 68 test cases passed.
 */
func FindMinHeightTrees(n int, edges [][]int) []int {
	list := make(map[int][]int)
	for _, edge := range edges {
		source := edge[0]
		destination := edge[1]
		neighbourList, contains := list[source]
		setConnectivity(&contains, &destination, &list, &source, &neighbourList)
		neighbourList, contains = list[destination]
		setConnectivity(&contains, &source, &list, &destination, &neighbourList)
	}

	distances := []int{}
	cache := make(map[int]int)
	for source:= 0; source < n; source++{
		value, contains := cache[source]
		if !contains {
			visited := make([]bool, n)
			maxHeight := doDFSFromSource(source, &list, 0, &visited, &cache)
			distances = append(distances, maxHeight)
		} else {
			distances = append(distances, value)
		}
	}

	min := distances[0]
	for _, value := range distances {
		if value < min {
			min = value
		}
	}

	finalResult := []int{}

	for index, value := range distances {
		if value == min {
			finalResult = append(finalResult, index)
		}
	}
	return finalResult
}

func doDFSFromSource(sourceNode int, list *map[int][]int, totalHeight int, visited *[]bool, cache *map[int]int) int {
	value, contains := (*cache)[sourceNode]
	if !contains {
		if !(*visited)[sourceNode] {
			(*visited)[sourceNode] = true
			maxHeight := totalHeight
			for _, neighbour := range (*list)[sourceNode] {
				if !(*visited)[neighbour] {
					maxHeight = getMax(maxHeight, doDFSFromSource(neighbour, list, totalHeight+1, visited, cache))
				}
			}
			return maxHeight
		}
		(*cache)[sourceNode] = totalHeight - 1
		return totalHeight - 1
	} else {
		return value
	}
}

func setConnectivity(contains *bool, destination *int, list *map[int][]int, source *int, neighbourList *[]int) {
	if !*contains {
		neighbourList := []int{}
		neighbourList = append(neighbourList, *destination)
		(*list)[*source] = neighbourList
	} else {
		*neighbourList = append(*neighbourList, *destination)
		(*list)[*source] = *neighbourList
	}
}

func getMax(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}