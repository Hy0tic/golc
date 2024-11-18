package main

// leetcode #3
func lengthOfLongestSubstring(s string) int {
    l := 0
	r := 1

	maxLen := 0

	for r < len(s) + 1 {
		substring := s[l : r]
		strLen := len(substring)

		// if there are no repeating character
		if strLen == uniqueCharacterCount(substring) { 
			maxLen = max(maxLen, strLen)
		} else {
			l++
		}
		r++
	}

	return maxLen
}

func uniqueCharacterCount(s string) int {
	// Create a map to track unique characters
	charMap := make(map[rune]bool)

	// Loop through each character in the string
	for _, char := range s {
		charMap[char] = true // Mark the character as seen
	}

	// The number of unique characters is the size of the map
	return len(charMap)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}