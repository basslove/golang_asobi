package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	text1 := "あ"
	for i := 0; i < len(text1); i++ {
		b := text1[i]
		fmt.Println(b)        // byte
		fmt.Printf("%x\n", b) // 16進数(utf8: E3 81 82)
	}

	fmt.Println("---------------------------------------")

	text2 := "あ"
	for _, v := range text2 {
		//　ひとつの文字は複数byteで表現される可能性があり、文字を表すbyteをまとめて読まないと正しい文字として認識できない
		// その解消として、文字を数える単位としてはbyteではなくcode pointのほうが都合がいい
		// runeはcode pointを提供する
		fmt.Println(v)         // rune
		fmt.Println(string(v)) // マルチバイト文字
	}

	text3 := "あいう"
	fmt.Println([]rune(text3))
	fmt.Println([]byte(text3))

	text4 := "いや〜ん"
	fmt.Println(len(text4))                    // byte数
	fmt.Println(len([]rune(text4)))            // rune数
	fmt.Println(utf8.RuneCountInString(text4)) // rune数

	text5 := "ここがオイラのこと"
	fmt.Println(text5[9:18])
	fmt.Println(string([]rune(text5)[3:6]))

	text6 := "ここがアチキのこと"
	byteIdx := strings.IndexRune(text6, 'ア')
	runeIdx := len([]rune(text6[0:byteIdx])) + 1
	fmt.Println(runeIdx)

	text7 := "ワッチのことを何卒よろしくお願いします"
	fmt.Println(strings.IndexRune(text7, '卒'))
	fmt.Println(strings.Contains(text7, "何卒"))
	fmt.Println(strings.ContainsRune(text7, '何'))

}
