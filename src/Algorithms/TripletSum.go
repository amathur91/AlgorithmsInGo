package algorithms

import "fmt"
/**
Solution to Triplet Sum Problem
 */
//func main(){
//	arr := []int{12, 3, 4, 1, 6, 9}
//	target := 24
//	findTriplet(arr, target)
//
//
//	//Test Case 2
//	arr2 := []int{1, 2, 3, 4, 5}
//	target2 := 9
//	findTriplet(arr2, target2)
//}

func findTriplet(arr []int, target int) {
	numMap := make(map[int]bool)
	numMap[arr[0]] = true
	for firstIndex := 1; firstIndex < len(arr)-1; firstIndex++ {
		for secondIndex := firstIndex + 1; secondIndex < len(arr); secondIndex++ {
			sum := arr[firstIndex] + arr[secondIndex]
			requiredNum := target - sum
			_, contains := numMap[requiredNum]
			if contains {
				fmt.Printf("Found the number : %d,%d,%d ", arr[firstIndex], arr[secondIndex], requiredNum)
				fmt.Println()
			}
		}
		numMap[arr[firstIndex]] = true
	}
}
