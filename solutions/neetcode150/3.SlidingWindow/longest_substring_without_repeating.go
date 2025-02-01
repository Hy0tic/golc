package main

// leetcode #3
func lengthOfLongestSubstring(s string) int {
	l := 0
	r := 1
	maxLen := 0
	strMap := make(map[byte]bool)

	for r < len(s) {
		// while substring contains duplicate, move L up
		for strMap[s[r]] {
			strMap[s[l]] = false
			l++
		}

		// substring no longer contain duplicate/substring does not contain duplicate
		strMap[s[r]] = true
		maxLen = Max(maxLen, r-l+1)
		r++
	}

	return maxLen
}
