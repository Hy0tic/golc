package main

// leetcode #424
func characterReplacement(s string, k int) int {
    left := 0
    maxFreq := 0
    result := 0
    freqMap := make(map[byte]int)

    for right := 0; right < len(s); right++ {
        freqMap[s[right]]++
        maxFreq = Max(maxFreq, freqMap[s[right]])

        // Shrink the window if the number of replacements exceeds k
        for right-left+1-maxFreq > k {
            freqMap[s[left]]--
            left++
        }

        result = max(result, right-left+1)
    }

    return result
}