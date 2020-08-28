package memorydatastore

var dataRepository = make(map[string]IndexRepository)

func saveToRepository(indexName *string, processedData *map[string][]DataFrequency) {
	data := IndexRepository{processedData: processedData}
	dataRepository[*indexName] = data
}

func getFromRepository(indexName string) *map[string][]DataFrequency {
	value, contains :=  dataRepository[indexName]
	if contains {
		return value.processedData
	}
	return nil
}


