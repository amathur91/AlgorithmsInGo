package memorydatastore

import "strings"

func processData(dataRequest *MemoryStoreDataRequest) {
	wordIndexMap := make(map[string][]DataFrequency)
	dataSet := dataRequest.DataSet
	for index, data := range dataSet {
		words := strings.Split(data, " ")
		for _, word := range words {
			 _,contains := wordIndexMap[word]
			if contains {
				//check if this index is already present
				dataFrequencyForWord := wordIndexMap[word]
				returnIndex := checkIfIndexAlreadyPresent(&index, &dataFrequencyForWord)
				if returnIndex == nil{
					wordIndexMap[word] = append(wordIndexMap[word], DataFrequency{index: index, frequency: 1, data: data})
				} else {
					wordIndexMap[word][*returnIndex].frequency++
				}
			} else {
				var dataFrequencies []DataFrequency
				dataFrequencies = append(dataFrequencies, DataFrequency{index: index, frequency: 1,data: data})
				wordIndexMap[word] = dataFrequencies
			}
		}
	}
	//we have a map which can tell us that for word apple how many string contain it and frequency in that particular string
	saveToRepository(&(dataRequest.IndexName), &wordIndexMap)
}

func checkIfIndexAlreadyPresent(targetIndex *int, listOfIndexes *[]DataFrequency) *int {
	for index, data := range *listOfIndexes {
		if data.index == *targetIndex {
			return &index
		}
	}
	return nil
}
