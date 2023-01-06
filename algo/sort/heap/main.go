package main

import (
	"fmt"
	"math/rand"
	"time"
)

func maxHeap(numbers []int, position int) {
	size := len(numbers)
	max := position
	leftChild := 2*position + 1
	rightChild := leftChild + 1
	if leftChild < size && numbers[leftChild] > numbers[position] {
		max = leftChild
	}
	if rightChild < size && numbers[rightChild] > numbers[max] {
		max = rightChild
	}

	if position != max {
		numbers[position], numbers[max] = numbers[max], numbers[position]
		maxHeap(numbers, max)
	}
}

func heap(numbers []int) {
	for i := (len(numbers) - 1) / 2; i >= 0; i-- {
		maxHeap(numbers, i)
	}
}

func heapSort(numbers []int) {
	heap(numbers)
	for i := len(numbers) - 1; i >= 1; i-- {
		numbers[i], numbers[0] = numbers[0], numbers[i]
		maxHeap(numbers[:i-1], 0)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	nums := make([]int, 0)
	for range make([]int, 5) {
		nums = append(nums, rand.Intn(100))
	}

	fmt.Println(nums)
	heapSort(nums)
	fmt.Println(nums)
}
