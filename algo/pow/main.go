package main

import "fmt"

// 累乗
func pow(x float64, y int) float64 {
	if y == 0 {
		return 1
	} else if z := pow(x, y/2); y%2 == 0 {
		return z * z
	} else {
		return x * z * z
	}
}

func main() {
	fmt.Println(pow(2, 16))
	fmt.Println(pow(2, 32))
	fmt.Println(pow(2, 64))
}
