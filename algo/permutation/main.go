package main

import "fmt"

// 順列
func permutation(n int, perm []int, used []bool) {
	if n == len(perm) {
		fmt.Println(perm)
	} else {
		for i := 1; i <= len(perm); i++ {
			if used[i] {
				continue
			}
			perm[n] = i
			used[i] = true
			permutation(n+1, perm, used)
			used[i] = false
		}
	}
}

func main() {
	permutation(0, make([]int, 4), make([]bool, 5))
}
