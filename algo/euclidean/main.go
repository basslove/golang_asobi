package main

import "fmt"

// 最大公約数
func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

// 最小公倍数
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	fmt.Println(gcd(42, 30))
	fmt.Println(gcd(15, 70))
	fmt.Println(lcm(5, 7))
	fmt.Println(lcm(14, 35))
}
