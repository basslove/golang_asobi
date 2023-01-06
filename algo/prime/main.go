package main

import "fmt"

func isPrime(target, primeSize int, nums []int) bool {
	for i := 1; i < primeSize; i++ {
		num := nums[i]
		if num*num > target {
			break
		}
		if target%num == 0 {
			return false
		}
	}
	return true
}

func main() {
	nums := make([]int, 100)
	nums[0] = 2
	primeSize := 1 // 素数の個数

	for n := 3; n < len(nums); n += 2 {
		if isPrime(n, primeSize, nums) {
			nums[primeSize] = n
			primeSize++
		}
	}
	fmt.Println(nums[:primeSize])
}
