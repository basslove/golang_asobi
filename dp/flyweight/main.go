package main

import (
	"fmt"
	"strconv"
)

type bigChar struct {
	charName string
	fontData string
}

func NewBigChar(charName string) *bigChar {
	char := &bigChar{charName: charName}
	char.loadFontData()
	return char
}

func (bc *bigChar) loadFontData() {
	num, _ := strconv.Atoi(bc.charName)
	var str string
	for i := 0; i < num; i++ {
		str += "-"
	}
	bc.fontData = str
}

func (bc *bigChar) Print() string {
	return bc.fontData
}

var instance *bigCharFactory

type bigCharFactory struct {
	pool map[string]*bigChar
}

func GetBigCharFactory() *bigCharFactory {
	if instance == nil {
		instance = &bigCharFactory{}
	}
	return instance
}

func (bcf *bigCharFactory) getBigChar(charName string) *bigChar {
	if bcf.pool == nil {
		bcf.pool = make(map[string]*bigChar)
	}
	if _, ok := bcf.pool[charName]; !ok {
		bcf.pool[charName] = NewBigChar(charName)
	}
	return bcf.pool[charName]
}

type bigString struct {
	bigChars []*bigChar
}

func NewBigString(str string) *bigString {
	bigStr := &bigString{}
	factory := GetBigCharFactory()

	for _, s := range str {
		bigStr.bigChars = append(bigStr.bigChars, factory.getBigChar(string(s)))
	}
	return bigStr
}

func (bs *bigString) Print() string {
	var result string
	for _, s := range bs.bigChars {
		result += s.Print() + "\n"
	}
	return result
}

func main() {
	fmt.Println("flyweight")

	bigStr := NewBigString("123456789")
	fmt.Println(bigStr.Print())
}
