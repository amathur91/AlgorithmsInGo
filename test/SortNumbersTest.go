package main

import (
	algorithms "Algorithms"
	"fmt"
)

func main() {
	case1()
	input := []int{2,0,1}
	algorithms.SortNumbers(&input)
	fmt.Println(input)
}

func case1() {
	input := []int{2, 0, 2, 1, 1, 0}
	algorithms.SortNumbers(&input)
	fmt.Println(input)
}