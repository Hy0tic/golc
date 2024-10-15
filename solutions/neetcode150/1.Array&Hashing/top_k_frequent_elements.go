package main

import "sort"

// Leetcode #347
func topKFrequent(nums []int, k int) []int {
	countMap := map[int]int{} // map of value and occurence
	result := []int{}

	for _, n := range nums {
		countMap[n]++
	}

	// convert map to array of arrays where the first value is num and second is num count
	var countArr [][]int
	for i, num := range countMap {
		countArr = append(countArr, []int{i, num})
	}

	sort.Slice(countArr, func(i, j int) bool {
		return countArr[i][1] > countArr[j][1]
	}) // sort arrays by num count AKA second array value

	// take the top k elements
	for i := 0; i < k; i++ {
		result = append(result, countArr[i][0])
	}

	return result
}
