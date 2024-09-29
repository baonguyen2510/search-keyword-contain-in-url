package common

func RemoveDuplicatesArrMap(results []map[string]interface{}) []map[string]interface{} {
	seenURLs := make(map[string]bool)
	uniqueResults := []map[string]interface{}{}

	for _, result := range results {
		url, ok := result["url"].(string)
		if !ok {
			continue
		}

		if !seenURLs[url] {
			uniqueResults = append(uniqueResults, result)
			seenURLs[url] = true
		}
	}

	return uniqueResults
}
