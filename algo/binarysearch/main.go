package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func binarySearch(target int, nums []int) bool {
	low := 0
	high := len(nums) - 1

	for low <= high {
		middle := low + (high-low)/2
		if nums[middle] == target {
			return true
		} else if nums[middle] < target {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}

	return false
}

func main() {
	fmt.Println("binary search")
	rand.Seed(time.Now().UnixNano())

	nums := make([]int, 0)
	for range make([]int, 5) {
		nums = append(nums, rand.Intn(100))
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	fmt.Println(binarySearch(nums[3], nums))
}
