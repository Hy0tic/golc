package main

// leetcode #242
func isAnagram(s string, t string) bool {
	map1 := map[byte]int{}
	map2 := map[byte]int{}

	for i := 0; i < len(s); i++ {
		map1[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		map2[t[i]]++
	}

	// Iterate through the first map
	for key, value1 := range map1 {
		if value2, exists := map2[key]; exists {
			// Compare values from both maps
			if value1 == value2 {
				continue
			} else {
				return false
			}
		} else {
			return false
		}
	}

	for key := range map2 {
		if _, exists := map1[key]; !exists {
			return false
		}
	}

	return true
}
