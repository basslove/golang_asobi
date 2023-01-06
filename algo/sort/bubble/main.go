package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(numbers []int) []int {
	lenNumbers := len(numbers)

	for i := range make([]int, lenNumbers) {
		for j := range make([]int, lenNumbers-1-i) {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	return numbers
}

func main() {
	rand.Seed(time.Now().UnixNano())

	nums := make([]int, 0)
	for range make([]int, 5) {
		nums = append(nums, rand.Intn(100))
	}

	fmt.Println(nums)
	fmt.Println(bubbleSort(nums))
}
