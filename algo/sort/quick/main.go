package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(low, high int, numbers []int) {
	pivot := numbers[low+(high-low)/2]
	i, j := low, high
	for {
		for pivot > numbers[i] {
			i++
		}
		for pivot < numbers[j] {
			j--
		}
		if i >= j {
			break
		}
		numbers[i], numbers[j] = numbers[j], numbers[i]
		i++
		j--
	}
	if low < i-1 {
		quickSort(low, i-1, numbers)
	}
	if high > j+1 {
		quickSort(j+1, high, numbers)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	nums := make([]int, 0)
	for range make([]int, 5) {
		nums = append(nums, rand.Intn(100))
	}

	fmt.Println(nums)
	quickSort(0, 4, nums)
	fmt.Println(nums)
}
