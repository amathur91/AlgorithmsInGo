package memorydatastore

func processQuery(query *string, indexName *string) []string{
	wordMap := getFromRepository(*indexName)
	data, contains := (*wordMap)[*query]
	if !contains  {
		return nil
	} else {
		return transformData(data)
	}
}
