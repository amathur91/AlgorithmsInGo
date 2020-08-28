package memorydatastore

// The payload that user will send for storing in the memory store
type MemoryStoreDataRequest struct {
	DataSet []string
	IndexName string
}

// The payload that the user will send for searching in datastore
type QueryRequest struct {
	IndexName string
	Keyword string
}

// The response payload that the MemoryDataStore will return with search results
type QueryResponse struct {
	searchResults []string
}

// This will the data structure that we will eventually store in memory
type ProcessedData struct {
	data string
	frequency int
}

type DataFrequency struct {
	index int
	frequency int
	data string
}

type IndexRepository struct {
	processedData *map[string][]DataFrequency
}

