package memorydatastore

//this map tell us that do we have data for the index
var indexMap map[string]bool = make(map[string]bool)


func PutData(dataStoreRequest *MemoryStoreDataRequest) bool {
	//call the validator to validate the request if that is fine then we
	//call the processor to process and return us the results
	if validateRequest(dataStoreRequest) {
		indexMap[dataStoreRequest.IndexName] = true
		processData(dataStoreRequest)
		return true
	}
	return false
}

func SearchData(queryRequest *QueryRequest) []string {
	//add check for query validation
	_, contains := indexMap[queryRequest.IndexName]
	if !contains {
		return nil
	} else {
		return processQuery(&queryRequest.Keyword, &queryRequest.IndexName)
	}
}
