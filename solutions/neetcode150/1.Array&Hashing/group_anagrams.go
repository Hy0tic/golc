package main

import (
	"encoding/json"
	"log"
)

// leetcode #49
func groupAnagrams(strs []string) [][]string {
	result := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		letterMap := make(map[byte]int)
		currentStr := strs[i];

		alphabet := "abcdefghijklmnopqrstuvwxyz"
		for k := 0; k < len(alphabet); k++ {
			letterMap[alphabet[k]] = 0
		}

        for k := 0; k < len(currentStr); k++ {
			letterMap[currentStr[k]]++;
		}

		// serialize letter map representing the string
		jsonString, err := json.Marshal(letterMap)
		if err != nil {
			log.Fatalf("Error serializing map: %v", err)
		}

		result[string(jsonString)] = append(result[string(jsonString)], currentStr)
    }


	var arrayOfArrays [][]string
	// Iterate over the map and append each slice to the array of arrays
	for _, values := range result {
		arrayOfArrays = append(arrayOfArrays, values)
	}
	
	return arrayOfArrays
}