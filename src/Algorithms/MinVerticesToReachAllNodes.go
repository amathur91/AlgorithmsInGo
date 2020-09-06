package algorithms

import "sort"
/**
https://leetcode.com/problems/minimum-number-of-vertices-to-reach-all-nodes/submissions/

Result:

61 / 66 test cases passed.
Status: Time Limit Exceeded


 */
func FindSmallestSetOfVertices(n int, edges [][]int) []int {
	// create a matrix for storing relations
	connMatrix := make([][]int, n)
	for index, _ := range connMatrix {
		connMatrix[index] = make([]int, n)
	}
	for _, edge := range edges {
		connMatrix[edge[0]][edge[1]] = 1
	}

	connCount := make(map[int][]int)

	cache := make(map[int][]int)
	for nodeIndex := 0; nodeIndex < n; nodeIndex++ {
		visited := make([]bool, n)
		listOfNodes := make(map[int]bool)
		value, contains := cache[nodeIndex]
		if !contains {
			doDFS (nodeIndex, &connMatrix, &visited, &listOfNodes, &cache)
			list := []int{}
			for key := range listOfNodes{
				list = append(list, key)
			}
			cache[nodeIndex] = list
			connCount[nodeIndex] = list
		} else {
			connCount[nodeIndex] = value
		}
	}

	//sort the conncount based on the size
	keys := []int{}
	for key := range connCount {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return len(connCount[keys[i]]) > len(connCount[keys[j]])
	})
	if len(connCount[keys[0]]) == n {
		return []int{keys[0]}
	}
	finalResult := findMinKeysFortheGraph(keys, connCount, n,)
	return finalResult
}

func findMinKeysFortheGraph(keys []int, matrix map[int][]int, totalNodes int) []int {
	set := make(map[int]bool)
	visited := make(map[int]bool)
	for _, value := range keys {
		_, contains := visited[value]
		if !contains {
			set[value] = true
			visited[value] = true
			neighbours := matrix[value]
			for _, neighbour := range neighbours {
				visited[neighbour] = true
			}
		}
	}
	finalAnswer := []int{}
	for key := range set {
		finalAnswer = append(finalAnswer, key)
	}

	return finalAnswer
}

func doDFS(index int, matrix *[][]int, visited *[]bool, listOfNodes *map[int]bool, cache *map[int][]int) {
	value, contains := (*cache)[index]
	if !contains {
		if !(*visited)[index] {
			(*visited)[index] = true
			(*listOfNodes)[index] = true
			for index, connected := range (*matrix)[index] {
				if connected == 1 {
					doDFS(index, matrix, visited, listOfNodes, cache)
				}
			}
		}
	} else {
		for _, node := range value {
			(*listOfNodes)[node] = true
		}

	}
}