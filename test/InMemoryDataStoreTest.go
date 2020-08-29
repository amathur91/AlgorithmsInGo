package main

import (
	"fmt"
	. "memorydatastore"
)
func main() {
	data := []string {"Apple Apple is good for health",
					"Orange is great for eyes",
					"I like Apple Products"}
	memoryDataStoreRequest := MemoryStoreDataRequest{IndexName: "NewIndex", DataSet: data}
	result := PutData(&memoryDataStoreRequest)
	fmt.Println("Saving data in memoryStore : ", result)
	searchResult := SearchData(&QueryRequest{IndexName: "NewIndex", Keyword: "Apple"})
	fmt.Println("Search Result : ")
	for _, value := range(searchResult) {
		fmt.Println(value)
	}
}
