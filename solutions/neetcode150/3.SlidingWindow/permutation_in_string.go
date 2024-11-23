package main


// leetcode #567
func checkInclusion(s1 string, s2 string) bool {
    if len(s2) < len(s1) {
        return false
    }

    // Initialize frequency maps for s1 and the current window in s2
    s1Freq := make([]int, 26)
    s2Freq := make([]int, 26)

    for i := 0; i < len(s1); i++ {
        s1Freq[s1[i]-'a']++
        s2Freq[s2[i]-'a']++
    }

    matches := 0
    for i := 0; i < 26; i++ {
        if s1Freq[i] == s2Freq[i] {
            matches++
        }
    }

    left := 0
    for right := len(s1); right < len(s2); right++ {
        if matches == 26 {
            return true
        }

        // Add the next character to the window
        index := s2[right] - 'a'
        s2Freq[index]++
        if s1Freq[index] == s2Freq[index] {
            matches++
        } else if s1Freq[index]+1 == s2Freq[index] {
            matches--
        }

        // Remove the left character from the window
        index = s2[left] - 'a'
        s2Freq[index]--
        if s1Freq[index] == s2Freq[index] {
            matches++
        } else if s1Freq[index]-1 == s2Freq[index] {
            matches--
        }
        left++
    }

    return matches == 26
}
