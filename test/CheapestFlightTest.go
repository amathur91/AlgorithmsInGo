package main

import (
	algorithms "Algorithms"
	"fmt"
)

func main() {
	n := 10
	edges :=[][]int{{3,4,4},{2,5,6},{4,7,10},{9,6,5},{7,4,4},{6,2,10},{6,8,6},{7,9,4},{1,5,4},{1,0,4},{9,7,3},{7,0,5},{6,5,8},{1,7,6},{4,0,9},{5,9,1},{8,7,3},{1,2,6},{4,1,5},{5,2,4},{1,9,1},{7,8,10},{0,4,2},{7,2,8}}
	src := 6
	dest := 0
	k := 7
	fmt.Println(algorithms.FindCheapestPrice(n, edges, src, dest, k))
}
