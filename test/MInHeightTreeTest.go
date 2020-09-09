package main

import (
	algorithms "Algorithms"
	"fmt"
)

func main() {
	//fmt.Println(algorithms.FindMinHeightTrees(4, [][]int{{1,0}, {1,2}, {1,3}}))
	fmt.Println(algorithms.FindMinHeightTrees(6, [][]int{{3,0},{3,1},{3,2},{3,4},{5,4}}))
}
