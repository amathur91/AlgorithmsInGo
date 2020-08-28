package memorydatastore

func validateRequest(request *MemoryStoreDataRequest) bool {
	if request != nil {
		if request.IndexName == "" || len(request.DataSet) == 0 {
			return false
		}
		return true
	}
	return false
}
