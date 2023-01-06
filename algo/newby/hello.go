package main

import "fmt"

// - hey
// - staticchek
// - go vet

func main() {
	fmt.Println("aaaaaaaa")
	var x []int // xはnilに初期化される
	fmt.Println(x == nil)
	x = append(x, 10)
	fmt.Println(x)

	var hoge = []int{1, 2, 3}
	// hoge = append(hoge, 4, 5, 6, 7)
	hoge2 := []int{20, 30, 40}
	hoge = append(hoge, hoge2...)
	fmt.Println(hoge)

	m := map[uint]int{1: 2}
	fmt.Println(m)

	var person struct { // 変数personを無名の構造体として
		name string
		age  int
		pet  string
	}
	person.name = "ボブ"
	person.age = 50
	person.pet = "dog"
	pet := struct {
		name string
		kind string
	}{
		name: "ポチ",
		kind: "dog",
	}
	fmt.Println(pet)
	fmt.Println(person)

	samples := []string{"hello", "apple_π!", "これは漢字文字列"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}

	var i any      // 「var i interface{} 」も可 i = 20
	fmt.Println(i) // 20
	i = "hello"
	fmt.Println(i) // hello
	i = struct {
		FirstName string
		LastName  string
	}{"信玄", "武田"}
	fmt.Println(i) // {信玄 武田}
}
