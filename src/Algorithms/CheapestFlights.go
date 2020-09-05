package algorithms

/**
Problem : https://leetcode.com/problems/cheapest-flights-within-k-stops/
Result : Runtime: 376 ms, faster than 18.18% of Go online submissions for Cheapest Flights Within K Stops.
Memory Usage: 16.2 MB, less than 6.06% of Go online submissions for Cheapest Flights Within K Stops.

47 / 47 test cases passed.
Status: Accepted
Runtime: 376 ms
Memory Usage: 16.2 MB

 */
var MAXINT = 217315

type CacheKey struct {
	src int
	dest int
	hopCount int
}

func FindCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	connMatrix := make([][]int, n)
	for index, _ := range connMatrix {
		connMatrix[index] = make([]int, n)
	}

	costMatrix := make([][]int, n)
	for index, _ := range costMatrix {
		costMatrix[index] = make([]int, n)
	}

	for _, flight := range flights {
		source := flight[0]
		dest := flight[1]
		cost := flight[2]
		connMatrix[source][dest] = 1
		costMatrix[source][dest] = cost
	}
	visited := make([]bool, n)
	cache := make(map[CacheKey]int)
	finalResult := findCheapPrice(&connMatrix, &costMatrix, &src, &dst, &K, 0, &visited, &cache)
	if finalResult == MAXINT {
		return -1
	}
	return finalResult
}

func findCheapPrice(connectivityMatrix *[][]int, costMatrix *[][]int, src *int, dest *int, hopsAllowed *int, currentHopCount int, visited *[]bool, cache *map[CacheKey]int) int {
	if currentHopCount >= *hopsAllowed {
		if (*connectivityMatrix)[*src][*dest] == 1 {
			return (*costMatrix)[*src][*dest]
		} else {
			return MAXINT
		}
	} else {
		//explore the neighbours
		minCost := MAXINT
		if (*connectivityMatrix)[*src][*dest] == 1 {
			minCost = getMin((*costMatrix)[*src][*dest], minCost)
		}

		for neighbourIndex := 0; neighbourIndex < len((*costMatrix)[*src]); neighbourIndex++ {
			if (*connectivityMatrix)[*src][neighbourIndex] == 1 && !(*visited)[neighbourIndex] {
				(*visited)[neighbourIndex] = true
				key := CacheKey{src: *src, dest: neighbourIndex, hopCount: currentHopCount+1}
				value, contains := (*cache)[key]
				if !contains {
					costFromNeighbour := findCheapPrice(connectivityMatrix, costMatrix, &neighbourIndex, dest, hopsAllowed, currentHopCount+1, visited, cache)
					(*cache)[key] = costFromNeighbour
					minCost = getMin(minCost, (*costMatrix)[*src][neighbourIndex]+ costFromNeighbour)
				} else {
					minCost = getMin(minCost, (*costMatrix)[*src][neighbourIndex]+ value)
				}
				(*visited)[neighbourIndex] = false
			}
		}
		return minCost
	}
}

func getMin(num1 int, num2 int) int {
	if num1 > num2 {
		return num2
	}
	return num1
}
