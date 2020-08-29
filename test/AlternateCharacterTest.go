package main

import (
	algorithms "Algorithms"
	"fmt"
)

func main() {
	testCase1()
	testCase2()
}

func testCase2() {
	input := "Geeks for geeks"
	result := algorithms.GetAlternateCharacterString(&input)
	fmt.Printf("Alternate Character String : %s \n", result)
}

func testCase1() {
	input := "It is a long day Dear."
	result := algorithms.GetAlternateCharacterString(&input)
	fmt.Printf("Alternate Character String : %s \n", result)
}
