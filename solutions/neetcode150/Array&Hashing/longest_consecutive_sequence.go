package main

func longestConsecutive(nums []int) int {
	var existMap = map[int]bool{}
	var longest = 0

	for i := 0; i < len(nums); i++ {
		existMap[nums[i]] = true
	}

	for i := 0; i < len(nums); i++ {
		// check if map contains nums[i] - 1
		var n = nums[i]

		if !existMap[n-1] {
			var counter = 0
			for existMap[n + counter] {
				counter++
			}

			if counter > longest {
				longest = counter
			}
		}

	}

    return longest
}