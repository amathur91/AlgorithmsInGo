package algorithms

//func main() {
//	arr1 := []int{1, 5, 9, 10, 15, 20}
//	arr2 := []int{2, 3, 8, 13}
//	sortArrays(&arr1, &arr2)
//	fmt.Println(arr1)
//	fmt.Println(arr2)
//}

func sortArrays(arr1 *[]int, arr2 *[]int) {
	startIndex := len(*arr2) - 1
	for index:= startIndex; index >=0; index-- {
		lastPositionIndex := len(*arr1) - 1
		lastPositionValue := (*arr1)[lastPositionIndex]
		for firstArrayPointer := len(*arr1) - 2; firstArrayPointer >=0; firstArrayPointer-- {
			if (*arr1)[firstArrayPointer] > (*arr2)[index] {
				//move it to the right
				(*arr1)[firstArrayPointer + 1] = (*arr1)[firstArrayPointer]
			} else {
				//place the current value from array 2 here and last position value here
				(*arr1)[firstArrayPointer + 1] = (*arr2)[index]
				(*arr2)[index] = lastPositionValue
				break
			}
		}
	}
}
