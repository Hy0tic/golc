package main

import (
    "sort"
)

type Key struct {
    A int
    B int
    C int
}

func threeSum(nums []int) [][]int {
    var hashset = map[Key]bool{}

    sort.Ints(nums)

    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        l := i + 1
        r := len(nums) - 1

        for l < r {
            var threeSum = nums[i] + nums[l] + nums[r]

            if threeSum > 0 {
                r--
            } else if threeSum < 0 {
                l++
            } else {
                hashset[Key { nums[i], nums[l], nums[r] }] = true
                l++

                for nums[l] == nums[l-1] && l < r {
                    l++
                }
            }
        }
    }

    res := [][]int{}

    for key, _ := range hashset {
        item := []int{ key.A, key.B, key.C }

        res = append(res, item)
    }

    return res
}