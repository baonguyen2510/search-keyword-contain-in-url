package common

import (
	"strconv"
)

func GetIntFromStr(s string) int64 {
	rs, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return rs
}

func GetFloatFromStr(s string) float64 {
	rs, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return rs
}

func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

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
