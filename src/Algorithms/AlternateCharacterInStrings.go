package algorithms

import "strings"

func GetAlternateCharacterString(input *string) string {
	resultString := ""
	characterUsageStore := make(map[string]bool)
	for index := 0; index < len(*input); index++ {
		character := strings.ToLower(string((*input)[index]))
		if isMarked(&character, characterUsageStore) {
			resultString += character
		}
	}
	return resultString
}

func isMarked(character *string, usageStore map[string]bool) bool {
	_ , contains := usageStore[*character]
	if !contains {
		//first time so add it to the map
		usageStore[*character] = true
		return true
	} else {
		delete(usageStore, *character)
		return false
	}
}

