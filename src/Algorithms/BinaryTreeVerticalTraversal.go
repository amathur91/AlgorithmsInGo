package main

import (
	"fmt"
	"models"
	"sort"
)

type nodelist struct {
	values []int
}

//func main() {
//	// Test case 1
//	rootNode := models.TreeNode{Val: 3}
//	rootNode.Left = &models.TreeNode{Val: 9}
//	rootNode.Right = &models.TreeNode{Val: 20}
//	rootNode.Right.Left = &models.TreeNode{Val: 15}
//	rootNode.Right.Right = &models.TreeNode{Val: 7}
//
//	// Test Case 2
//	printInvertedTree(&rootNode)
//}

func printInvertedTree (rootNode *models.TreeNode) ([][]int){
	var finalResult [][]int
	positionMap := make(map[int]nodelist)
	//populate the map
	traverseTree(rootNode, 0, positionMap)
	//print the map
	keys := make([]int, len(positionMap))
	position := 0

	//create a data store for the keys
	for index := range positionMap {
		keys[position] = index
		position++
	}
	// sorting the keys
	sort.Ints(keys)
	// traversing the map after sorting keys
	for _, key := range keys {
		nodes := positionMap[key]
		finalResult = append(finalResult, nodes.values)
	}
	fmt.Println(finalResult)
	return finalResult
}

func traverseTree(rootNode *models.TreeNode, verticalPosition int, positionMap map[int]nodelist) {
	if rootNode != nil {
		_, contains := positionMap[verticalPosition]
		if !contains {
			//create a list
			var datalist []int
			newDataList := append(datalist, rootNode.Val)
			positionMap[verticalPosition] = nodelist{values: newDataList}
		} else {
			updatedSlice := append(positionMap[verticalPosition].values, rootNode.Val)
			positionMap[verticalPosition] = nodelist{values: updatedSlice}
		}
		traverseTree(rootNode.Left, verticalPosition - 1, positionMap)
		traverseTree(rootNode.Right, verticalPosition + 1, positionMap)
	}
}
