package main

import (
	"fmt"
	"math/rand"
	"time"
)

func func_o1(numbers []int) {
	fmt.Println(numbers)
	fmt.Printf("result : %v", numbers)
}

func func_o_log_n(number int) {
	fmt.Println(number)

	if number <= 1 {
		fmt.Printf("result : %v", number)
	} else {
		fmt.Printf("process : %v", number)
		func_o_log_n(number / 2)
	}
}

func func_o_n(numbers []int) {
	fmt.Println(numbers)

	for idx, n := range numbers {
		fmt.Printf("%v time number: %v\n", idx+1, n)
	}
	fmt.Printf("result: %v time\n", len(numbers))
}

func func_o_n_log_n(number int) {
	fmt.Println(number)

	numbers := make([]int, 5)
	for idx, n := range numbers {
		fmt.Printf("%v time number: %v\n", idx+1, n)
	}

	if number <= 1 {
		fmt.Printf("result: %v\n", number)
		return
	}
	func_o_n_log_n(number / 2)
}

func func_o_n_square(numbers []int) {
	fmt.Println(numbers)

	count := make([]int, len(numbers))

	for i := range count {
		for j := range count {
			fmt.Printf("numbers[%v] : %v, numbers[%v] : %v\n", i, numbers[i], j, numbers[j])
		}
	}
}

func output() {
	rand.Seed(time.Now().Unix())

	fmt.Println("#### O(1) ####")
	func_o1(rand.Perm(10))
	fmt.Println()

	fmt.Println("#### O(log n)) ####")
	func_o_log_n(rand.Intn(10))
	fmt.Println()

	fmt.Println("#### O(n) ####")
	func_o_n(rand.Perm(10))
	fmt.Println()

	fmt.Println("#### O(n log n) ####")
	func_o_n_log_n(rand.Intn(10))
	fmt.Println()

	fmt.Println("#### O(n^2) ####")
	func_o_n_square(rand.Perm(10))
	fmt.Println()
}

func main() {
	output()
}
