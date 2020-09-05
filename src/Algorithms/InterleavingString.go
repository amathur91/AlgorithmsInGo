package algorithms

import (
	"strings"
)

/**
Problem : https://leetcode.com/problems/interleaving-string/
Result : Runtime: 0 ms, faster than 100.00% of Go online submissions for Interleaving String.
Memory Usage: 2.2 MB, less than 40.43% of Go online submissions for Interleaving String.

101 / 101 test cases passed.
Status: Accepted
Runtime: 0 ms
Memory Usage: 2.2 MB

 */
type CacheKey struct {
	firstPtr int
	secondPtr int
	thirdPtr int
}

func IsInterleave(s1 string, s2 string, s3 string) bool {
	//sanity Check
	if len(s1) + len(s2) != len(s3) {
		return false
	}
	firstStrPtr := 0
	firstStrLength := len(s1)
	secondStringPtr := 0
	secondStrLength := len(s2)
	cacheMap := make(map[CacheKey]bool)
	return checkStringInterleaving(&s3, firstStrPtr, firstStrLength, secondStringPtr, secondStrLength, &s1, &s2, 0, &cacheMap)

}

func checkStringInterleaving(targetString *string, firstStrPtr int, firstStringLength int, secondStrPtr int, secondStrLength int, firstString *string, secondString *string, thirdStringPtr int, cache *map[CacheKey]bool) bool {
	if firstStrPtr == firstStringLength  && secondStrPtr == secondStrLength && thirdStringPtr == len(*targetString) {
		return true //we have reached length of the target string matching all the characters of 1 and 2nd string
	}
	var result1 = false
	var result2 = false
	var letterAtFirstString  = ""
	var letterAtSecondString  = ""
	letterAtThirdString := string((*targetString)[thirdStringPtr])
	if firstStrPtr < firstStringLength {
		letterAtFirstString = string((*firstString)[firstStrPtr])
		if strings.EqualFold(letterAtFirstString, letterAtThirdString) {
			key := CacheKey{firstPtr: firstStrPtr + 1, secondPtr: secondStrPtr, thirdPtr: thirdStringPtr + 1}
			value, contains := (*cache)[key]
			if !contains {
				result1 = checkStringInterleaving(targetString, firstStrPtr+1, firstStringLength, secondStrPtr, secondStrLength, firstString, secondString, thirdStringPtr+1, cache)
				(*cache)[key] = result1
			} else {
				result1 = value
			}
		}
	}
	if secondStrPtr < secondStrLength {
		letterAtSecondString = string((*secondString)[secondStrPtr])
		if strings.EqualFold(letterAtSecondString, letterAtThirdString) {
			key := CacheKey{firstPtr: firstStrPtr, secondPtr: secondStrPtr + 1, thirdPtr: thirdStringPtr + 1}
			value, contains := (*cache)[key]
			if !contains {
				result2 = checkStringInterleaving(targetString, firstStrPtr, firstStringLength, secondStrPtr+1, secondStrLength, firstString, secondString, thirdStringPtr+1, cache)
			} else {
				result2 = value
			}
		}
	}

	if !strings.EqualFold(letterAtSecondString, letterAtThirdString) && !strings.EqualFold(letterAtFirstString, letterAtThirdString) {
		//if the character at third string does not match either of the first and second character
		//then return false
		return false
	}
	return result1 || result2
}

