package algorithms

/*Leetcode: https://leetcode.com/problems/word-subsets/
Results :
Runtime: 256 ms, faster than 16.67% of Go online submissions for Word Subsets.
Memory Usage: 7.6 MB, less than 72.22% of Go online submissions for Word Subsets.


55 / 55 test cases passed.
Status: Accepted
Runtime: 256 ms
Memory Usage: 7.6 MB
*/

//func main() {
//	groupA := []string{"amazon","apple","facebook","google","leetcode"}
//	groupB := []string{"ec","oc","ceo"}
//	result := wordSubsets(groupA, groupB)
//	fmt.Println(result)
//}
func wordSubsets(groupA []string, groupB []string) []string {
	characterMap := make(map[string]int)
	buildMap(&characterMap, &groupB)
	finalResult := make([]string, 0, 10000)
	for _, word := range groupA {
		localMap := buildLocalMap(&word)
		result := true
		for key := range characterMap {
			value, contains := (*localMap)[key]
			if !contains {
				result = false
			} else {
				if value < characterMap[key] {
					result = false
				}
			}
		}
		if result {
			finalResult = append(finalResult, word)
		}
	}
	return finalResult
}

func buildMap(characterMap *map[string]int, words *[]string) {
	for _, word := range *words {
		localMap := buildLocalMap(&word)
		//compare and add this to global map
		for key := range *localMap {
			_, contains := (*characterMap)[key]
			if !contains {
				(*characterMap)[key] =(*localMap)[key]
			} else {
				if (*localMap)[key] > (*characterMap)[key] {
					(*characterMap)[key] = (*localMap)[key]
				}
			}
		}
	}
}

func buildLocalMap(word *string) *map[string]int {
	localMap := make(map[string]int)
	for index := 0; index < len(*word); index++ {
		character := string( (*word)[index])
		_, contains := localMap[character]
		if !contains {
			localMap[character] = 1
		} else {
			currentValue := localMap[character]
			localMap[character] = currentValue + 1
		}
	}
	return &localMap
}


