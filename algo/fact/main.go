package main

import (
	"fmt"
)

// 再帰
func fact(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * fact(n-1)
	}
}

// 末尾再帰
func facti(n, a int) int {
	if n == 0 {
		return a
	} else {
		return facti(n-1, a*n)
	}
}

func main() {
	for i := range make([]int, 30) {
		fmt.Println(i, ":", fact(i))
	}

	for i := range make([]int, 30) {
		fmt.Println(i, ":", facti(i, 1))
	}
}
