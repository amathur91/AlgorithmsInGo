package main

import (
	"sort"
	"strings"
)
/**
Leetcode : https://leetcode.com/problems/top-k-frequent-words/
Level : Medium
Result : Runtime: 12 ms, faster than 16.67% of Go online submissions for Top K Frequent Words.
Memory Usage: 4.5 MB, less than 49.41% of Go online submissions for Top K Frequent Words.
 */
type wordDataStore struct {
	word string
	count int
}
//func main() {
//	input := []string{"i", "love", "leetcode", "i", "love", "coding"}
//	k := 2
//	result := getTopKFrequentWords(&input, k)
//	fmt.Println(result)
//
//	input2 := []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
//	k2 := 4
//	result2 := getTopKFrequentWords(&input2, k2)
//	fmt.Println(result2)
//}

func getTopKFrequentWords(input *[]string, k int) []string {
	//build a word frequency count
	wordFrequencyMap := make(map[string]int)
	for _, word := range *input {
		_, contains := wordFrequencyMap[word]
		if !contains {
			wordFrequencyMap[word] = 1
		} else {
			currentCount := wordFrequencyMap[word]
			wordFrequencyMap[word] = currentCount + 1
		}
	}

	//traverse map and create a slice of data
	var wordSlice []wordDataStore
	for key := range  wordFrequencyMap {
		count := wordFrequencyMap[key]
		wordData := wordDataStore{word: key, count: count}
		wordSlice = append(wordSlice, wordData)
	}
	sort.Slice(wordSlice, func(i, j int) bool {
		if  wordSlice[i].count !=  wordSlice[j].count {
			return wordSlice[i].count > wordSlice[j].count
		} else {
			result :=  strings.Compare(wordSlice[i].word, wordSlice[j].word)
			if result == -1 {
				return true
			} else {
				return false
			}
		}
	})

	words := make([]string, 0, k)

	count := 0
	for _, value := range wordSlice {
		if count == k {
			return words
		}
		words = append(words, value.word)
		count++
	}
	return words
}
