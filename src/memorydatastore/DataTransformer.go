package memorydatastore

//Implement sort later. That is TODO functionality.
func transformData(dataFrequency []DataFrequency) []string {
	var returnData []string
	for _, value := range dataFrequency {
		returnData = append(returnData, value.data)
	}
	return returnData
}
